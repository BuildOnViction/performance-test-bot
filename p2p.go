package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
)

const messageId = 0

type Message string

func MyProtocol() p2p.Protocol {
	return p2p.Protocol{
		Name:    "MyProtocol",
		Version: 1,
		Length:  1,
		Run:     msgHandler,
	}
}

func startServer() *p2p.Server {
	key, _ := crypto.GenerateKey()
	var peers []*discover.Node
	node, _ := discover.ParseNode("enode://11ad11a89e90e7bd9aaea10fe4cead6cea0122a5827068d1d442dd1f779f303fd9cd9cbf04ee1a8301354d1449efd9cc818e68eeb67765dc363d18f7fae7d1e1@127.0.0.1:30304")
	peers = append(peers, node)
	config := p2p.Config{
		Name:           "Bot",
		MaxPeers:       10,
		ListenAddr:     ":30303",
		PrivateKey:     key,
		Protocols:      []p2p.Protocol{MyProtocol()},
		BootstrapNodes: peers,
	}
	server := &p2p.Server{
		Config: config,
	}
	return server
}

func msgHandler(peer *p2p.Peer, ws p2p.MsgReadWriter) error {
	for {
		msg, err := ws.ReadMsg() // 3.
		if err != nil {          // 4.
			return err // if reading fails return err which will disconnect the peer.
		}

		var myMessage [1]Message
		err = msg.Decode(&myMessage) // 5.
		if err != nil {
			// handle decode error
			continue
		}

		switch myMessage[0] {
		case "foo":
			err := p2p.SendItems(ws, messageId, "bar") // 6.
			if err != nil {
				return err // return (and disconnect) error if writing fails.
			}
		default:
			fmt.Println("recv:", myMessage)
		}
	}

	return nil
}
