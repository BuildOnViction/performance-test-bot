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

const key = `{"address":"f2851114a967578f9dba4a92a3f7aa09237e4d08","crypto":{"cipher":"aes-128-ctr","ciphertext":"5b2b7158195c753e3f29e4db0474d765499a73855fd5c96bf3a836c97086cff9","cipherparams":{"iv":"5562bdf087417f3c730f58ca0ef11357"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"2f3c0240758dab124b40ec7d1fc42be02e853be655f51aad3b2fc2c6678aed5a"},"mac":"0eb6dc95cc3f92595d88b7cea507929fb576b1c89afe61e9090c4900a6471d69"},"id":"098ebf16-0cd8-4cd6-aab8-1a48eb93168c","version":3}`

var nonce = uint64(5600)

func Sender(url string) {
	fmt.Println(url)
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client: %v", err)
	}
	fmt.Println("OK!!!!!")

	rawTransaction(client)
	wg.Done()
}

func rawTransaction(client *ethclient.Client) {
	d := time.Now().Add(100000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	unlockedKey, err := keystore.DecryptKey([]byte(key), "")
	nonce = nonce + 1

	if err != nil {
		fmt.Println("Wrong passcode")
	} else {
		tx := types.NewTransaction(nonce, common.HexToAddress("0x56724a9e4d2bb2dca01999acade2e88a92b11a9e"), big.NewInt(12400000), 21000, big.NewInt(1000000000), nil)
		signTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(40686)), unlockedKey.PrivateKey)
		fmt.Println(signTx)
		err = client.SendTransaction(ctx, signTx)

		fmt.Println(err, nonce)
	}
}
