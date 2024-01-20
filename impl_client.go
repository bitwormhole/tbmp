package tbmp

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

type clientImpl struct {
	name string
}

func (inst *clientImpl) _impl() Client {
	return inst
}

func (inst *clientImpl) Name() string {
	return inst.name
}

func (inst *clientImpl) SetName(name string) {
	inst.name = name
}

func (inst *clientImpl) prepareConfig(aURL string, options []Option) (*Configuration, error) {

	u, err := url.Parse(aURL)
	if err != nil {
		return nil, err
	}

	cfg := new(Configuration)
	for _, op := range options {
		op(cfg)
	}

	host := u.Hostname()
	portStr := u.Port()
	port, _ := strconv.Atoi(portStr)

	if cfg.Host == "" {
		cfg.Host = host
	}

	if cfg.Port < 1 {
		cfg.Port = port
	}

	if cfg.Protocol == "" {
		cfg.Protocol = NameProtocol
	}

	return cfg, nil
}

func (inst *clientImpl) prepareHeaders(url string, cfg *Configuration) {

	name := inst.name
	method := cfg.Method
	protocol := cfg.Protocol

	if name == "" {
		name = NameClient
	}

	if method == "" {
		method = http.MethodGet
	}

	if protocol == "" {
		protocol = NameProtocol
	}

	h := &cfg.Headers
	h.Set(HeaderMethod, method)
	h.Set(HeaderURL, url)
	h.Set(HeaderClient, name)
	h.Set(HeaderProtocol, protocol)
}

func (inst *clientImpl) Connect(url string, options ...Option) (ClientSideConnection, error) {

	cfg, err := inst.prepareConfig(url, options)
	if err != nil {
		return nil, err
	}

	inst.prepareHeaders(url, cfg)
	hlist := &cfg.Headers

	conn1, err := inst.openConn(cfg)
	if err != nil {
		return nil, err
	}

	conn2 := new(clientConnImpl)
	defer func() {
		if conn1 != nil {
			conn1.Close()
		}
	}()

	cc := &connectionContext{
		config: cfg,
		conn:   conn1,
		closer: conn1,
		reader: conn1,
		writer: conn1,
	}

	csc, err := conn2.init(cc)
	if err != nil {
		return nil, err
	}

	_, err = conn2.Upstream(hlist)
	if err != nil {
		return nil, err
	}

	_, err = conn2.Downstream()
	if err != nil {
		return nil, err
	}

	conn1 = nil
	return csc, nil
}

func (inst *clientImpl) openConn(opt *Configuration) (net.Conn, error) {
	if opt.Secure {
		return inst.openAsTLS(opt)
	}
	return inst.openAsPlain(opt)
}

func (inst *clientImpl) openAsPlain(opt *Configuration) (net.Conn, error) {
	network := "tcp"
	addr := opt.getAddress()
	addr2, err := net.ResolveTCPAddr(network, addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP(network, nil, addr2)
	return conn, err
}

func (inst *clientImpl) openAsTLS(opt *Configuration) (net.Conn, error) {
	cfg := &tls.Config{}
	network := "tcp"
	addr := opt.getAddress()
	conn, err := tls.Dial(network, addr, cfg)
	return conn, err
}

////////////////////////////////////////////////////////////////////////////////

type clientConnRx struct {
	cc *connectionContext
}

func (inst *clientConnRx) _impl() RxStream {
	return inst
}

func (inst *clientConnRx) Headers() *Headers {
	return &inst.cc.rxHdr
}

func (inst *clientConnRx) Read(buf []byte) (int, error) {
	return inst.cc.reader.Read(buf)
}

////////////////////////////////////////////////////////////////////////////////

type clientConnTx struct {
	cc *connectionContext
}

func (inst *clientConnTx) _impl() TxStream {
	return inst
}

func (inst *clientConnTx) Headers() *Headers {
	return &inst.cc.txHdr
}

func (inst *clientConnTx) Write(buf []byte) (int, error) {
	return inst.cc.writer.Write(buf)
}

////////////////////////////////////////////////////////////////////////////////

type clientConnImpl struct {
	cc *connectionContext
	rx RxStream
	tx TxStream
}

func (inst *clientConnImpl) _impl() ClientSideConnection {
	return inst
}

func (inst *clientConnImpl) init(cc *connectionContext) (ClientSideConnection, error) {

	inst.cc = cc
	inst.rx = &clientConnRx{cc: cc}
	inst.tx = &clientConnTx{cc: cc}

	return inst, nil
}

func (inst *clientConnImpl) Close() error {
	return inst.cc.Close()
}

func (inst *clientConnImpl) Downstream() (RxStream, error) {

	if !inst.cc.rxHdrDone {
		err := inst.cc.receiveHead()
		if err != nil {
			return nil, err
		}
	}

	return inst.rx, nil
}

func (inst *clientConnImpl) Upstream(src ...*Headers) (TxStream, error) {

	dst := &inst.cc.txHdr
	for _, hlist := range src {
		hlist.ForEach(func(name, value string) {
			dst.Set(name, value)
		})
	}

	if !inst.cc.txHdrDone {
		err := inst.cc.sendHead()
		if err != nil {
			return nil, err
		}
	}

	return inst.tx, nil
}
