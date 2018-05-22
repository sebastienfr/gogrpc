//go:generate protoc -I ../protocol --go_out=plugins=grpc:../protocol ../protocol/spirits.proto

package main

import (
	"context"
	"github.com/satori/go.uuid"
	sfr "github.com/sebastienfr/gogrpc/protocol"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// spiritHandler is used to implement SpiritServiceServer
type spiritHandler struct{}

func (sh *spiritHandler) CreateSpirit(ctx context.Context, request *sfr.SpiritCreationRequest) (*sfr.SpiritCreationResponse, error) {
	log.Printf("New spirit creation received %+v", request)
	spirit := request.GetSpirit()
	spirit.Id = uuid.NewV4().String()
	response := &sfr.SpiritCreationResponse{
		Spirit: spirit,
	}
	return response, nil
}

func (sh *spiritHandler) UpdateSpirit(ctx context.Context, request *sfr.SpiritUpdateRequest) (*sfr.SpiritUpdateResponse, error) {
	log.Printf("New spirit update received %+v", request)
	return nil, nil
}

func (sh *spiritHandler) DeleteSpirit(ctx context.Context, request *sfr.SpiritDeleteRequest) (*sfr.SpiritDeleteResponse, error) {
	log.Printf("New spirit deletion received %+v", request)
	return nil, nil
}

func (sh *spiritHandler) SearchSpirit(ctx context.Context, request *sfr.SpiritSearchRequest) (*sfr.SpiritSearchResponse, error) {
	log.Printf("New spirit search received %+v", request)
	return nil, nil
}

func main() {
	// create server socket
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// build new spirit server
	s := grpc.NewServer()
	sfr.RegisterSpiritServiceServer(s, &spiritHandler{})

	// register reflection service on gRPC spiritHandler.
	//reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
