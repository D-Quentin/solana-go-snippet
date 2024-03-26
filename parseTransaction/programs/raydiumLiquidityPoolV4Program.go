package programs

import (
	"solana-poc/parseTransaction/types"
)

var swapBaseInArgs = []*types.Args{{Name: "AmountIn", Type: "uint64"}, {Name: "MinimumAmountOut", Type: "uint64"}}
var swapBaseOutArgs = []*types.Args{{Name: "maxAmountIn", Type: "uint64"}, {Name: "amountOut", Type: "uint64"}}

func LoadRaydiumLiquidityPoolV4Program() map[byte]types.ProgramHandlerDecoder {
	decoder := make(map[byte]types.ProgramHandlerDecoder)
	decoder[9] = func(data []byte) ([]*types.Args, string, error) {
		return DecodeArgs(swapBaseInArgs, data), "SwapBaseIn", nil
	}
	decoder[10] = func(data []byte) ([]*types.Args, string, error) {
		return DecodeArgs(swapBaseOutArgs, data), "SwapBaseOut", nil
	}
	return decoder
}
