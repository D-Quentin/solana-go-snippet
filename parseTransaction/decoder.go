package main

import (
	"errors"
	"slices"
	"solana-poc/parseTransaction/programs"
	"solana-poc/parseTransaction/types"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var tokenProgramID = solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
var computeBudgetProgramID = solana.MustPublicKeyFromBase58("ComputeBudget111111111111111111111111111111")
var RaydiumLiquidityPoolV4ProgramID = solana.MustPublicKeyFromBase58("675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8")

func createDecoder() types.Decoder {
	decoder := make(map[solana.PublicKey]map[byte]types.ProgramHandlerDecoder)
	decoder[tokenProgramID] = programs.LoadTokenProgram()
	decoder[computeBudgetProgramID] = programs.LoadComputeBudget()
	decoder[RaydiumLiquidityPoolV4ProgramID] = programs.LoadRaydiumLiquidityPoolV4Program()
	return decoder
}

func innerInstructionsDecoder(innerInstructions rpc.InnerInstruction, accountKeys []solana.PublicKey, decoder types.Decoder) []*types.DecodedInstruction {
	var decodedInstructions []*types.DecodedInstruction

	for _, innerInstruction := range innerInstructions.Instructions {
		programID := accountKeys[innerInstruction.ProgramIDIndex]
		newInstruction := &types.DecodedInstruction{
			ProgramID: programID,
		}
		if _, ok := decoder[programID]; ok {
			instructionHandlerIndex := innerInstruction.Data[0]
			newInstruction.InstructionHandlerIndex = instructionHandlerIndex
			if _, ok := decoder[programID][instructionHandlerIndex]; ok {
				decodedArgs, handlerName, err := decoder[programID][instructionHandlerIndex](innerInstruction.Data)
				if err != nil {
					newInstruction.Error = err
				}
				newInstruction.Args = decodedArgs
				newInstruction.InstructionHandlerName = handlerName
			} else {
				newInstruction.Error = errors.New(("unknown instruction handler"))
			}
		} else {
			newInstruction.Error = errors.New("unknown program ID")
		}
		decodedInstructions = append(decodedInstructions, newInstruction)
	}
	return decodedInstructions
}

func InstructionsDecoder(instructions []solana.CompiledInstruction, accountKeys []solana.PublicKey, allInnerInstructions []rpc.InnerInstruction) []*types.DecodedInstruction {
	decoder := createDecoder()
	var decodedInstruction []*types.DecodedInstruction

	for i, instruction := range instructions {
		programID := accountKeys[instruction.ProgramIDIndex]
		innerInstructionIndex := slices.IndexFunc(allInnerInstructions, func(innerInstruction rpc.InnerInstruction) bool { return innerInstruction.Index == uint16(i) })
		newInstruction := &types.DecodedInstruction{
			ProgramID: programID,
		}
		if innerInstructionIndex != -1 {
			newInstruction.InnerInstructions = innerInstructionsDecoder(allInnerInstructions[innerInstructionIndex], accountKeys, decoder)
		}
		if _, ok := decoder[programID]; ok {
			instructionHandlerIndex := instruction.Data[0]
			newInstruction.InstructionHandlerIndex = instructionHandlerIndex
			if _, ok := decoder[programID][instructionHandlerIndex]; ok {
				decodedArgs, handlerName, err := decoder[programID][instructionHandlerIndex](instruction.Data)
				if err != nil {
					newInstruction.Error = err
				}
				newInstruction.Args = decodedArgs
				newInstruction.InstructionHandlerName = handlerName
			} else {
				newInstruction.Error = errors.New(("unknown instruction handler"))
			}
		} else {
			newInstruction.Error = errors.New("unknown program ID")
		}
		decodedInstruction = append(decodedInstruction, newInstruction)
	}
	return decodedInstruction
}
