
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

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	machine.RegisterMachineServer(s, &MachineServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func testExecute_Live(t *testing.T, client machine.MachineClient, instructions []*machine.Instruction, wants []float32) {
	log.Printf("Streaming %v", instructions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.Execute(ctx)
	if err != nil {
		log.Fatalf("%v.Execute(ctx) = %v, %v: ", client, stream, err)
	}
	waitc := make(chan struct{})
	go func() {
		i := 0
		for {
			result, err := stream.Recv()
			if err == io.EOF {
				log.Println("EOF")
				close(waitc)
				return
			}
			if err != nil {
				log.Printf("Err: %v", err)
			}
			log.Printf("output: %v", result.GetOutput())
			got := result.GetOutput()
			want := wants[i]
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
			i++
		}
	}()

	for _, instruction := range instructions {
		if err := stream.Send(instruction); err != nil {
			log.Fatalf("%v.Send(%v) = %v: ", stream, instruction, err)
		}