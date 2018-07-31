package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"

	"github.com/silverspase/blockchain/blockchain"
	"github.com/silverspase/blockchain/server"
	"github.com/silverspase/blockchain/types"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := types.Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		blockchain.Blockchain = append(blockchain.Blockchain, genesisBlock)
	}()
	log.Fatal(server.Run())

}
