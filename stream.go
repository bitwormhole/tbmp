package tbmp

import "io"

// RxStream ...
type RxStream interface {
	io.Reader

	Headers() *Headers
}

// TxStream ...
type TxStream interface {
	io.Writer

	Headers() *Headers
}

// Upstream ...
type Upstream struct {
	// request

	Protocol    string
	ContentType string
	Headers     Headers

	Method   string
	URL      string
	Location Location
}

// Downstream ...
type Downstream struct {
	Protocol    string
	ContentType string
	Headers     Headers

	Status int
}
