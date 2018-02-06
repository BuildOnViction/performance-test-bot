package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
	"time"
)

var (
	NWorkers  = flag.Int("n", 4, "The number of workers to start")
	CUrl      = flag.String("url", "http://localhost:8545", "That you want to connect")
	NReq      = flag.Int("req", 1, "The number of transactions")
	BootNodes = flag.String("bootnodes", "", "Bootstrap nodes for peer to peer network")
	Port      = flag.Int("port", 30303, "Node port")
)

var wg sync.WaitGroup

const key = `{"address":"ba47474654eed2ba872678804611bb8cbe22016a","crypto":{"cipher":"aes-128-ctr","ciphertext":"7b6a3452555995ccc04abfea2e7a2bf95acf5098e2328c6a51994a12c617d2c3","cipherparams":{"iv":"9bc61d03b69151139803a727eab692df"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ec14069205247d65a6b71a4d6dd1a905bbd3eeb6f319dda548214cc5365f780d"},"mac":"b5c43a99dfcab6670e4b4c7d83cd596d803587ebbab422533170e51ec85fa235"},"id":"45262826-d26b-4bd4-9e00-526235a26aa6","version":3}`

var client *ethclient.Client
var nonce uint64
var unlockedKey *keystore.Key

func main() {
	flag.Parse()

	server := startServer()
	if err := server.Start(); err != nil {
		fmt.Println("Could not start server: %v", err)
	}
	fmt.Println("Server started", server.NodeInfo().Enode)

	client, _ = ethclient.Dial(*CUrl)
	d := time.Now().Add(100000 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	unlockedKey, _ = keystore.DecryptKey([]byte(key), "")
	nonce, _ = client.NonceAt(ctx, unlockedKey.Address, nil)

	wg.Add(*NReq)

	// Start the dispatcher.
	StartDispatcher(*NWorkers)

	for i := 0; i < *NReq; i++ {
		Collector(uint64(i) + nonce)
	}

	wg.Wait()

	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println(server.PeerCount())
		}
	}
}
