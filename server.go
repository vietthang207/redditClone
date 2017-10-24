package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"redditClone/databases"
)

type ServerConfig struct {
	Port      string `json: port`
	QuerySize int    `json: query_size`
}

func readConfigFile() ServerConfig {
	return ServerConfig{"8080", 20}
}

var (
	db *databases.Database
)

func main() {
	router := NewRouter()
	config := readConfigFile()
	db = databases.NewDatabase(config.QuerySize)
	f, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("Start to serve on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
