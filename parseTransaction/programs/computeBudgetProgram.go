package programs

import (
	"solana-poc/parseTransaction/types"
)

var setComputeUnitLimitArgs = []*types.Args{{Name: "Compute Unit Limit", Type: "uint32"}}
var setComputeUnitPriceArgs = []*types.Args{{Name: "Compute Unit Price", Type: "uint64"}}

func LoadComputeBudget() map[byte]types.ProgramHandlerDecoder {
	decoder := make(map[byte]types.ProgramHandlerDecoder)
	decoder[2] = func(data []byte) ([]*types.Args, string, error) {
		return DecodeArgs(setComputeUnitLimitArgs, data), "SetComputeUnitLimit", nil
	}
	decoder[3] = func(data []byte) ([]*types.Args, string, error) {
		return DecodeArgs(setComputeUnitPriceArgs, data), "SetComputeUnitPrice", nil
	}
	return decoder
}
