package main

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/lib/pq"
	recordapi "github.com/md3756/ustart_tutorial/api/grpc/record"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config recordapi.Config
	//Importing configuration from json
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	//Generating api object
	recordService, err := recordapi.New(&config)
	if err != nil {
		panic(err)
	}

	recordService.Run()
}
