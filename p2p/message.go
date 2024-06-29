package p2p

import "net"

// RPC holds any arbitrary message data that is being sent over each transport between two nodes
type RPC struct {
	From    net.Addr
	Payload []byte
}
