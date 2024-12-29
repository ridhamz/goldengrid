package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPTransport is a struct that represents a TCP connection.
type TCPPeer struct {
	// conn is the underlying TCP connection.
	conn net.Conn
	// Outbound is false if the connection is incoming, true if it is outgoing.
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder       Decoder
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    NOPHandshakeFunc,
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
			continue
		}
		go t.handleConn(conn)
	}
}

type Temp struct{}

// handle the connection
func (t *TCPTransport) handleConn(conn net.Conn) {
	//peer := NewTCPPeer(conn, true)
	if err := t.shakeHands(conn); err != nil {
		fmt.Printf("Handshake error: %s\n", err)
		conn.Close()
		return
	}

	// Read loop
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {

			fmt.Printf("Decode error: %s\n", err)
			continue
		}
	}
}
