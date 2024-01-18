package tbmp

import (
	"crypto/tls"
	"net"

	"github.com/starter-go/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type serverImpl struct {
	name string
}

func (inst *serverImpl) _impl() Server {
	return inst
}

func (inst *serverImpl) Name() string {
	return inst.name
}

func (inst *serverImpl) SetName(name string) {
	inst.name = name
}

func (inst *serverImpl) Listen(cfg *Configuration) (Listener, error) {
	if cfg == nil {
		cfg = new(Configuration)
	}
	if cfg.Secure {
		return inst.listenTLS(cfg)
	}
	return inst.listenPlain(cfg)
}

func (inst *serverImpl) listenPlain(opt *Configuration) (Listener, error) {
	// use TCP
	addr1 := opt.getAddress()
	network := "tcp"
	addr2, err := net.ResolveTCPAddr(network, addr1)
	if err != nil {
		return nil, err
	}
	vlog.Info("listen TBMP (over TCP) at %s", addr2.String())
	l, err := net.ListenTCP(network, addr2)
	if err != nil {
		return nil, err
	}
	facade := &listenerImpl{
		inner: l,
		name:  inst.name,
		cfg:   opt,
	}
	return facade, nil
}

func (inst *serverImpl) listenTLS(opt *Configuration) (Listener, error) {
	// use TLS
	addr := opt.getAddress()
	network := "tcp"
	cfg := &tls.Config{}
	vlog.Info("listen TBMP (over TLS) at %s", addr)
	l, err := tls.Listen(network, addr, cfg)
	if err != nil {
		return nil, err
	}
	facade := &listenerImpl{
		inner: l,
		name:  inst.name,
		cfg:   opt,
	}
	return facade, nil
}

////////////////////////////////////////////////////////////////////////////////

type listenerImpl struct {
	inner net.Listener
	cfg   *Configuration
	name  string
}

func (inst *listenerImpl) _impl() Listener {
	return inst
}

func (inst *listenerImpl) Close() error {
	return inst.inner.Close()
}

func (inst *listenerImpl) Accept() (ServerSideConnection, error) {

	conn1, err := inst.inner.Accept()
	if err != nil {
		return nil, err
	}

	cc := &connectionContext{
		config: inst.cfg,
		conn:   conn1,
		closer: conn1,
		reader: conn1,
		writer: conn1,
	}
	inst.prepareHeaders(cc, inst.cfg)

	conn2 := &serverConnImpl{}
	return conn2.init(cc)
}

func (inst *listenerImpl) prepareHeaders(cc *connectionContext, cfg *Configuration) {

	name := inst.name
	protocol := cfg.Protocol

	if name == "" {
		name = NameServer
	}

	if protocol == "" {
		protocol = NameProtocol
	}

	cc.txHdr.Set(HeaderProtocol, protocol)
	cc.txHdr.Set(HeaderServer, name)
}

////////////////////////////////////////////////////////////////////////////////

type serverConnTx struct {
	cc *connectionContext
}

func (inst *serverConnTx) _impl() TxStream {
	return inst
}

func (inst *serverConnTx) Headers() *Headers {
	return &inst.cc.txHdr
}

func (inst *serverConnTx) Write(b []byte) (int, error) {
	return inst.cc.writer.Write(b)
}

////////////////////////////////////////////////////////////////////////////////

type serverConnRx struct {
	cc *connectionContext
}

func (inst *serverConnRx) _impl() RxStream {
	return inst
}

func (inst *serverConnRx) Headers() *Headers {
	return &inst.cc.rxHdr
}

func (inst *serverConnRx) Read(b []byte) (int, error) {
	return inst.cc.reader.Read(b)
}

////////////////////////////////////////////////////////////////////////////////

type serverConnImpl struct {
	cc *connectionContext
	rx RxStream
	tx TxStream
}

func (inst *serverConnImpl) _impl() ServerSideConnection {
	return inst
}

func (inst *serverConnImpl) init(cc *connectionContext) (ServerSideConnection, error) {

	inst.cc = cc
	inst.rx = &serverConnRx{cc: cc}
	inst.tx = &serverConnTx{cc: cc}

	return inst, nil
}

func (inst *serverConnImpl) Close() error {
	return inst.cc.Close()
}

func (inst *serverConnImpl) Upstream() (RxStream, error) {

	if !inst.cc.rxHdrDone {
		err := inst.cc.receiveHead()
		if err != nil {
			return nil, err
		}
	}

	return inst.rx, nil
}

func (inst *serverConnImpl) Downstream(src ...*Headers) (TxStream, error) {

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
