package main

import (
	"fmt"
	"log"
	"net/http"
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
	fmt.Println("Start to serve on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
