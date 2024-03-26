package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gagliardetto/solana-go"
	lookup "github.com/gagliardetto/solana-go/programs/address-lookup-table"
	"github.com/gagliardetto/solana-go/rpc"
)

func PrettyPrint(data interface{}) {
	var p []byte
	p, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

const TRANSACTION_HASH string = "6KFQoLizxiFCa6HjPf2FVxySUzWdPz44PcSrHTgmx6BH8swgjgyDTHBF1x2Yp2rcPQT2SXgcLNxGR8mZRQ8bsXP" // Raydium swap transaction

var MAX_SUPPORTED_TRANSACTION_VERSION uint64 = 0

// Get all the missing address from the transaction with address lookups
func processTransactionWithAddressLookups(txx solana.Transaction, rpcClient *rpc.Client) solana.Transaction {
	tblKeys := txx.Message.GetAddressTableLookups().GetTableIDs()
	resolutions := make(map[solana.PublicKey]solana.PublicKeySlice)
	for _, key := range tblKeys {
		info, err := rpcClient.GetAccountInfo(
			context.Background(),
			key,
		)
		if err != nil {
			panic(err)
		}
		tableContent, err := lookup.DecodeAddressLookupTableState(info.GetBinary())
		if err != nil {
			panic(err)
		}
		resolutions[key] = tableContent.Addresses
	}
	err := txx.Message.SetAddressTables(resolutions)
	if err != nil {
		panic(err)
	}
	err = txx.Message.ResolveLookups()
	if err != nil {
		panic(err)
	}
	return txx
}

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
	tx, err := out.Transaction.GetTransaction()
	if err != nil {
		panic(err)
	}
	transaction := processTransactionWithAddressLookups(*tx, client)
	instructions := transaction.Message.Instructions
	accounts := transaction.Message.AccountKeys
	innerInstructions := out.Meta.InnerInstructions

	decodedInstructions := InstructionsDecoder(instructions, accounts, innerInstructions)

	PrettyPrint(decodedInstructions)
}
