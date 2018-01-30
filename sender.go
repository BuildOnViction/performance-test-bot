package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

func Sender(no uint64) {
	d := time.Now().Add(100000 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	tx := types.NewTransaction(no, common.HexToAddress("0x56724a9e4d2bb2dca01999acade2e88a92b11a9e"), big.NewInt(12400000), 21000, big.NewInt(1000000000), nil)
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(40686)), unlockedKey.PrivateKey)
	fmt.Println(signTx)
	err = client.SendTransaction(ctx, signTx)

	fmt.Println(err, no)
	wg.Done()
}
