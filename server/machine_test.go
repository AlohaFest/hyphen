
//
// machine_test.go
// Copyright (C) 2020 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package server

import (
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/toransahu/grpc-eg-go/machine"
	"github.com/toransahu/grpc-eg-go/mock_machine"
)

func TestExecute(t *testing.T) {
	s := MachineServer{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockServerStream := mock_machine.NewMockMachine_ExecuteServer(ctrl)