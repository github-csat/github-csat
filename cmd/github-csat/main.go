package main

import (
	"github.com/dexhorthy/github-csat/pkg/server"
	"log"
)

func main() {

	if err := server.Main(); err != nil {
		log.Fatal(err)
	}
}
