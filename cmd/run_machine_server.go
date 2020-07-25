
//
// run_machine_server.go
// Copyright (C) 2020 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/toransahu/grpc-eg-go/machine"
	"github.com/toransahu/grpc-eg-go/server"