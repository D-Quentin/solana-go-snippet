package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

var MAX_SUPPORTED_TRANSACTION_VERSION uint64 = 0

func main() {
	client, err := ws.Connect(context.Background(), "wss://private-node-url") // A private node ws connection with the flag --rpc-pubsub-enable-block-subscription
	if err != nil {
		panic(err)
	}

	sub, err := client.BlockSubscribe(ws.NewBlockSubscribeFilterAll(), &ws.BlockSubscribeOpts{
		MaxSupportedTransactionVersion: &MAX_SUPPORTED_TRANSACTION_VERSION, // Is needed to query recent blocks
	})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	for {
		got, err := sub.Recv()
		if err != nil {
			panic(err)
		}
		spew.Dump(got)
		/*
			Successful output:
			(*ws.BlockResult) {
				Context {
					Slot (uint64)
				}
				Value {
					Slot (uint64)
					Block (*rpc.GetBlockResult)
				}
			}
		*/
	}
}
