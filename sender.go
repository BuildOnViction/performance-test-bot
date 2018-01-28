package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

const key = `{"address":"ba47474654eed2ba872678804611bb8cbe22016a","crypto":{"cipher":"aes-128-ctr","ciphertext":"7b6a3452555995ccc04abfea2e7a2bf95acf5098e2328c6a51994a12c617d2c3","cipherparams":{"iv":"9bc61d03b69151139803a727eab692df"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ec14069205247d65a6b71a4d6dd1a905bbd3eeb6f319dda548214cc5365f780d"},"mac":"b5c43a99dfcab6670e4b4c7d83cd596d803587ebbab422533170e51ec85fa235"},"id":"45262826-d26b-4bd4-9e00-526235a26aa6","version":3}`

var url = "https://core.tomocoin.io"
var client, e = ethclient.Dial(url)
var d = time.Now().Add(100000 * time.Millisecond)
var ctx, er = context.WithDeadline(context.Background(), d)
var unlockedKey, err = keystore.DecryptKey([]byte(key), "")
var n, errr = client.NonceAt(ctx, unlockedKey.Address, nil)
var nonce = n - 1

//var nonce = uint64(1257)

func Sender(no uint64) {
	fmt.Println(url)
	//d := time.Now().Add(10000 * time.Millisecond)
	//ctx, cancel := context.WithDeadline(context.Background(), d)
	//defer cancel()
	//unlockedKey, err := keystore.DecryptKey([]byte(key), "")
	//nonce, _ := client.NonceAt(ctx, unlockedKey.Address, nil)
	//nonce = nonce + 1
	//nonce := uint64(958)

	if err != nil {
		fmt.Println("Wrong passcode")
	} else {
		tx := types.NewTransaction(no, common.HexToAddress("0x56724a9e4d2bb2dca01999acade2e88a92b11a9e"), big.NewInt(12400000), 21000, big.NewInt(1000000000), nil)
		signTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(40686)), unlockedKey.PrivateKey)
		fmt.Println(signTx)
		err = client.SendTransaction(ctx, signTx)

		fmt.Println(err, no)
	}
	wg.Done()
}
