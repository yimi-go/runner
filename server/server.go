package server

import (
	"errors"
	"net"

	"github.com/yimi-go/runner"
)

var errNotServing = errors.New("server: the server is not serving")

// IsErrNotServing determines if an error is server not serving error.
// It supports wrapped errors.
func IsErrNotServing(err error) bool {
	return errors.Is(err, errNotServing)
}

// NotServing produces an error that told the server is not serving.
func NotServing() error {
	return errNotServing
}

// Runner is a Runner which run a server.
type Runner interface {
	runner.Runner
	// Address returns the address the Server is serving on.
	// If the Server is not serving, an error returns.
	Address() (addr net.Addr, err error)
}
