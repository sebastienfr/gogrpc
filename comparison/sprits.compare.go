package main

import (
	"encoding/json"
	"github.com/gogo/protobuf/proto"
	"github.com/satori/go.uuid"
	"github.com/sebastienfr/gogrpc/protocol"
	"io/ioutil"
	"log"
	"time"
)

type SpiritJSON struct {
	Id           string
	Name         string
	Distiller    string
	Bottler      string
	Country      string
	Region       string
	Composition  string
	Type         gogrpc.Spirit_SpiritType
	Age          int32
	BottlingDate int64
	Score        float32
	Comment      string
}

func main() {
	sj := SpiritJSON{
		Age:          10,
		Id:           uuid.NewV4().String(),
		BottlingDate: time.Now().AddDate(-2, 0, 0).Unix(),
		Bottler:      "Seb",
		Comment:      "Hello Beer Drinkers !",
		Composition:  "Malt, Houblon",
		Country:      "France",
		Distiller:    "DevFest Lille",
		Name:         "Chti beer",
		Region:       "NPDC",
		Score:        10.0,
		Type:         gogrpc.Spirit_TypeBeer,
	}

	jsonb, err := json.Marshal(&sj)
	if err != nil {
		log.Fatalf("unable to marshal json spirit %+v", err)
	}

	err = ioutil.WriteFile("jsonbytes.json", jsonb, 0644)
	if err != nil {
		log.Fatalf("unable to write json spirit %+v", err)
	}

	scr := &gogrpc.SpiritCreationRequest{
		Spirit: &gogrpc.Spirit{
			Age:          10,
			Id:           uuid.NewV4().String(),
			BottlingDate: time.Now().AddDate(-2, 0, 0).Unix(),
			Bottler:      "Seb",
			Comment:      "Hello Beer Drinkers !",
			Composition:  "Malt, Houblon",
			Country:      "France",
			Distiller:    "DevFest Lille",
			Name:         "Chti beer",
			Region:       "NPDC",
			Score:        10.0,
			Type:         gogrpc.Spirit_TypeBeer,
		}}

	protobb, err := proto.Marshal(scr)
	if err != nil {
		log.Fatalf("unable to marshal protobuff spirit %+v", err)
	}

	err = ioutil.WriteFile("protobuffbytes.pb", protobb, 0644)
	if err != nil {
		log.Fatalf("unable to write protobuff spirit %+v", err)
	}

}
