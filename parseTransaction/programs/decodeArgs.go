package programs

import (
	"encoding/binary"
	"solana-poc/parseTransaction/types"
)

func DecodeArgs(args []*types.Args, data []byte) []*types.Args {
	decodedArgs := make([]*types.Args, len(args))
	for i, arg := range args {
		decodedArgs[i] = &types.Args{
			Name:  arg.Name,
			Type:  arg.Type,
			Value: arg.Value,
		}
	}

	currentByte := 1
	for _, arg := range decodedArgs {
		switch arg.Type {
		case "uint64":
			arg.Value = binary.LittleEndian.Uint64(data[currentByte : currentByte+8])
			currentByte += 8
		case "uint32":
			arg.Value = binary.LittleEndian.Uint32(data[currentByte : currentByte+4])
			currentByte += 4
		case "uint16":
			arg.Value = binary.LittleEndian.Uint16(data[currentByte : currentByte+2])
			currentByte += 2
		case "uint8":
			arg.Value = data[currentByte]
			currentByte++
		}
	}
	return decodedArgs
}
