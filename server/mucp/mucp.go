// Package mucp provides an mucp server
package mucp

import (
	"github.com/stack-labs/stack-rpc/server"
)

// NewServer returns a micro server interface
func NewServer(opts ...server.Option) server.Server {
	return server.NewServer(opts...)
}
