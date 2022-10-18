
//
// machine_mock_test.go
// Copyright (C) 2020 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package mock_machine_test

import (
	context "context"
	"log"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/toransahu/grpc-eg-go/machine"
	mock_machine "github.com/toransahu/grpc-eg-go/mock_machine"
)

func testExecute(t *testing.T, client machine.MachineClient) {