//go:generate protoc -I ../protocol --go_out=plugins=grpc:../protocol ../protocol/spirits.proto

package main

import (
	"log"
	"time"

	sfr "github.com/sebastienfr/gogrpc/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// build the spirit client
	c := sfr.NewSpiritServiceClient(conn)

	// create a new context for RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// do the the RPC
	r, err := c.CreateSpirit(ctx, &sfr.SpiritCreationRequest{
		Spirit: &sfr.Spirit{
			Age:          10,
			BottlingDate: time.Now().AddDate(-2, 0, 0).Unix(),
			Bottler:      "Seb",
			Comment:      "Hello Bier Drinkers !",
			Composition:  "Malt, Houblon",
			Country:      "France",
			Distiller:    "DevFest Lille",
			Name:         "Chti bier",
			Region:       "NPDC",
			Score:        10.0,
			Type:         sfr.Spirit_TypeBier,
		}})
	if err != nil {
		log.Fatalf("could not create: %v", err)
	}
	log.Printf("Created: %+v", r)
}
