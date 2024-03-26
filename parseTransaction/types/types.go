package types

import "github.com/gagliardetto/solana-go"

type Args struct {
	Name  string
	Type  string
	Value interface{}
}

type DecodedInstruction struct {
	ProgramID               solana.PublicKey
	InstructionHandlerName  string
	InstructionHandlerIndex byte
	Args                    []*Args
	Accounts                []solana.PublicKey
	InnerInstructions       []*DecodedInstruction
	Error                   error
}

type Decoder map[solana.PublicKey]map[byte]ProgramHandlerDecoder

type ProgramHandlerDecoder func(data []byte) ([]*Args, string, error)
