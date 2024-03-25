package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

const TRANSACTION_HASH string = "6KFQoLizxiFCa6HjPf2FVxySUzWdPz44PcSrHTgmx6BH8swgjgyDTHBF1x2Yp2rcPQT2SXgcLNxGR8mZRQ8bsXP" // Raydium swap transaction

var MAX_SUPPORTED_TRANSACTION_VERSION uint64 = 0

func main() {
	endpoint := rpc.MainNetBeta_RPC // Public Solana endpoint
	client := rpc.New(endpoint)
	txSig := solana.MustSignatureFromBase58(TRANSACTION_HASH)

	out, err := client.GetTransaction(
		context.TODO(),
		txSig,
		&rpc.GetTransactionOpts{
			MaxSupportedTransactionVersion: &MAX_SUPPORTED_TRANSACTION_VERSION, // Is needed to query recent blocks
		},
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)

	/*
		Successful output
		(*rpc.GetTransactionResult) {
			Slot (uint64),
			BlockTime (*solana.UnixTimeSeconds)
			Transaction (*rpc.TransactionResultEnvelope) {
				asDecodedBinary (solana.Data)
				asParsedTransaction (*solana.Transaction) {
					Signatures ([]solana.Signature)
					Message (*solana.Message) {
						Header (*solana.MessageHeader)
						AccountKeys ([]solana.PublicKey)
						RecentBlockhash (solana.Hash)
						Instructions ([]solana.CompiledInstruction)
					}
				}
			}
			Meta (*rpc.TransactionMeta) {
				Fee (uint64)
				PreBalances ([]uint64)
				PostBalances ([]uint64)
				InnerInstructions ([]rpc.TransactionInnerInstructions)
				PreTokenBalances ([]rpc.TokenBalance)
				PostTokenBalances ([]rpc.TokenBalance)
				LogMessages ([]string)
				Rewards ([]rpc.BlockReward)
				LoadedAddresses (LoadedAddresses) {
					ReadOnly ([]solana.PublicKey)
					Writable ([]solana.PublicKey)
				}
			}
			Version (rpc.TransactionVersion)
		}
	*/
}
