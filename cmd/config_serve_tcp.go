package cmd

import (
	"github.com/zrepl/zrepl/config"
	"time"
	"net"
)

type TCPListenerFactory struct {
	Address string
}

func parseTCPListenerFactory(c config.Global, in config.TCPServe) (*TCPListenerFactory, error) {

	lf := &TCPListenerFactory{
		Address: in.Listen,
	}
	return lf, nil
}

var TCPListenerHandshakeTimeout = 10 * time.Second // FIXME make configurable

func (f *TCPListenerFactory) Listen() (net.Listener, error) {
	return net.Listen("tcp", f.Address)
}