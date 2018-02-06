package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"strconv"
	"strings"
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
	bootNodes := strings.Split(*BootNodes, ",")

	var peers []*discover.Node
	for _, n := range bootNodes {
		if len(n) > 0 {
			node, _ := discover.ParseNode(n)
			peers = append(peers, node)
		}
	}
	config := p2p.Config{
		Name:           "Bot",
		MaxPeers:       10,
		ListenAddr:     ":" + strconv.Itoa(*Port),
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
