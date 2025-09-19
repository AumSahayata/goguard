package main

import (
	"log"

	"github.com/AumSahayata/goguard/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
