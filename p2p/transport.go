package p2p

// Peer is an interface that represents the remote node in the network.
type Peer interface {
}

// Transport is anything that handles the communication between nodes in the network.
// It could be a TCP connection, a UDP connection, or even a Unix socket.
type Transport interface {
	ListenAndAccept() error
}
