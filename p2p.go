package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"strconv"
	"strings"
	"time"
)

const messageId = 0

type Message struct {
	NReq int
}

func BotProtocol() p2p.Protocol {
	return p2p.Protocol{
		Name:    "BotProtocol",
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
		Protocols:      []p2p.Protocol{BotProtocol()},
		BootstrapNodes: peers,
	}
	server := &p2p.Server{
		Config: config,
	}
	return server
}

func msgHandler(peer *p2p.Peer, rw p2p.MsgReadWriter) error {
	for {
		select {
		case <-time.After(1 * time.Second):
			p2p.SendItems(rw, messageId, &Message{NReq: 1})
		}
		msg, err := rw.ReadMsg()
		if err != nil {
			return err
		}

		var myMessage Message
		err = msg.Decode(&myMessage)
		if err != nil {
			continue
		}

		switch myMessage.NReq {
		case 1:
			err := p2p.SendItems(rw, messageId, &Message{NReq: 2})
			if err != nil {
				return err
			}
		default:
			fmt.Println("recv:", myMessage)
		}
	}

	return nil
}
