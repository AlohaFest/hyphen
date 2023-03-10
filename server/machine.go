
//
// machine.go
// Copyright (C) 2020 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package server

import (
	"fmt"
	"io"
	"log"

	"github.com/toransahu/grpc-eg-go/machine"
	"github.com/toransahu/grpc-eg-go/utils"
	"github.com/toransahu/grpc-eg-go/utils/stack"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OperatorType string

const (
	PUSH OperatorType = "PUSH"
	POP               = "POP"
	ADD               = "ADD"
	SUB               = "SUB"
	MUL               = "MUL"
	DIV               = "DIV"
	FIB               = "FIB"
)

type MachineServer struct{}

// Execute runs the set of instructions given.
func (s *MachineServer) Execute(stream machine.Machine_ExecuteServer) error {
	var stack stack.Stack
	for {
		instruction, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF")
			return nil
		}
		if err != nil {
			return err
		}

		operand := instruction.GetOperand()
		operator := instruction.GetOperator()
		op_type := OperatorType(operator)

		fmt.Printf("Operand: %v, Operator: %v\n", operand, operator)

		switch op_type {
		case PUSH:
			stack.Push(float32(operand))
		case POP:
			stack.Pop()
		case ADD, SUB, MUL, DIV:
			item2, popped := stack.Pop()
			item1, popped := stack.Pop()

			if !popped {
				return status.Error(codes.Aborted, "Invalid sets of instructions. Execution aborted")
			}

			var res float32
			if op_type == ADD {
				res = item1 + item2
			} else if op_type == SUB {
				res = item1 - item2
			} else if op_type == MUL {
				res = item1 * item2
			} else if op_type == DIV {
				res = item1 / item2
			}

			stack.Push(res)
			if err := stream.Send(&machine.Result{Output: float32(res)}); err != nil {
				return err
			}
		case FIB:
			n, popped := stack.Pop()

			if !popped {
				return status.Error(codes.Aborted, "Invalid sets of instructions. Execution aborted")
			}

			if op_type == FIB {
				for f := range utils.FibonacciRange(int(n)) {
					if err := stream.Send(&machine.Result{Output: float32(f)}); err != nil {
						return err
					}
				}
			}
		default:
			return status.Errorf(codes.Unimplemented, "Operation '%s' not implemented yet", operator)
		}

	}
}