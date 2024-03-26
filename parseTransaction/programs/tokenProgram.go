package programs

import (
	"solana-poc/parseTransaction/types"
)

var transferArgs = []*types.Args{{Name: "Amount", Type: "uint64"}}

func LoadTokenProgram() map[byte]types.ProgramHandlerDecoder {
	decoder := make(map[byte]types.ProgramHandlerDecoder)
	decoder[3] = func(data []byte) ([]*types.Args, string, error) {
		return DecodeArgs(transferArgs, data), "Transfer", nil
	}
	return decoder
}
