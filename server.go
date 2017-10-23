package main

import (
	"fmt"
	"log"
	"net/http"
	"redditClone/routes"
)

type ServerConfig struct {
	Port string `json: port`
}

func readConfigFile() ServerConfig {
	return ServerConfig{"8080"}
}

func main() {
	router := routes.NewRouter()
	config := readConfigFile()
	fmt.Println("Start to serve on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
