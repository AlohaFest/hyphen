
This is basic example of gRPC based microservice to serve a Virtual Stack Machine over HTTP2.0 network.

# Pre-requisites
```bash
# install protobuf compiler
~/grpc-eg-go
$ go get -u -v github.com/golang/protobuf/{proto,protoc-gen-go}

# install grpc
~/grpc-eg-go
$ go get -u -v google.golang.org/grpc
```
# Generate Go Source Code from Protobuf
```bash
~/grpc-eg-go
$ SRC_DIR=./