package main

import (
	"log"

	"github.vom/ridhamz/goldengrid/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":5000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
