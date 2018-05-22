# GRPC Golang

The main purpose of this project is to get an overview about GRPC and how to use it with Go.

## Links

French slides associated to that repo are available here [disponibles ici](https://docs.google.com/presentation/d/1bOzQD5ytBqioqGNQk4hL9_smYrOD-l2205NFxqKL5uA/edit?usp=sharing)

## Setup for Go

It is supposed your environment is ready to compile [Go](https://golang.org/doc/install).   
To build [gRPC](https://grpc.io/docs/quickstart/go.html#go-version) you then need the following tools :
 * First [install](https://github.com/google/protobuf) protocol buffer compiler for your distribution.
 * Then [install](https://github.com/golang/protobuf) the Go protobuf extension.
 * Finally [install](https://grpc.io/docs/quickstart/go.html) gRPC for Go

## Live coding steps

### Defining protocol

### Building stub
    cd protocol
    protoc --go_out=plugins=grpc:. ./spirits.proto

### Integration in client and server

### Testing comparison