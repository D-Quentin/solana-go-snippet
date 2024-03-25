package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

const BLOCK_NUMBER uint64 = 1 // Block number to query

var MAX_SUPPORTED_TRANSACTION_VERSION uint64 = 0

func main() {
	endpoint := rpc.MainNetBeta_RPC // Public Solana endpoint
	client := rpc.New(endpoint)

	out, err := client.GetBlockWithOpts(context.TODO(), BLOCK_NUMBER, &rpc.GetBlockOpts{
		MaxSupportedTransactionVersion: &MAX_SUPPORTED_TRANSACTION_VERSION, // Is needed to query recent blocks
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
	/*
		Successful output:
		(*rpc.GetBlockResult) {
			Blockhash (solana.Hash)
			PrevBlockhash (solana.Hash)
			ParentSlot (uint64)
			Transactions ([]rpc.TransactionWithMeta)
			Signatures ([]solana.Signature)
			Rewards ([]rpc.BlockReward)
			BlockTime (*solana.UnixTimeSeconds)
			BlockHeight (*uint64)
		}
	*/
}
