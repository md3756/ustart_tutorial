package main

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/lib/pq"
	carapi "github.com/md3756/ustart_tutorial/api/grpc/car"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config carapi.Config
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
	carService, err := carapi.New(&config)
	if err != nil {
		panic(err)
	}

	carService.Run()
}
