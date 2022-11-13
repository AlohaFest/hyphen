
//
// machine_test.go
// Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package server

import (
	"context"
	"io"
	"log"
	"net"
	"testing"
	"time"

	"github.com/toransahu/grpc-eg-go/machine"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)
