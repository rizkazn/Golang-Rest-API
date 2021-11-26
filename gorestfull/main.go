package main

import (
	"log"
	"os"

	"github.com/rizkazn/gorestfull/configs/command"
)

func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal("Unable to Run App")
	}
}
