
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