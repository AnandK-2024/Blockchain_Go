package core

type OpCode byte

// 0x0 range - arithmetic ops.
const (
	PushInt    OpCode = 0x0
	ADD        OpCode = 0x1
	MUL        OpCode = 0x2
	SUB        OpCode = 0x3
	DIV        OpCode = 0x4
	SDIV       OpCode = 0x5
	MOD        OpCode = 0x6
	SMOD       OpCode = 0x7
	ADDMOD     OpCode = 0x8
	MULMOD     OpCode = 0x9
	EXP        OpCode = 0xa
	SIGNEXTEND OpCode = 0xb
	PushByte   OpCode = 0xc
	Pack       OpCode = 0xd
	Store      OpCode = 0xe
)
