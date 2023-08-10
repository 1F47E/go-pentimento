package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/1F47E/go-pentimento/cmd"
	"github.com/1F47E/go-pentimento/pkg/logger"
	"github.com/1F47E/go-pentimento/pkg/utils"
)

type Command string

const (
	CommandEncode Command = "encode"
	CommandDecode Command = "decode"
	CommandFit    Command = "fit"
)

func main() {
	log := logger.Log

	if len(os.Args) < 2 {
		utils.ExitWithManual()
	}

	// get command
	command := Command(os.Args[1])

	args := os.Args[2:]
	if command == CommandEncode {
		err := cmd.Encode(args)
		if err != nil {
			log.Fatalf("Error encoding image: %v", err)
		}
		utils.PrintGreen("Data encoded successfully")
	} else if command == CommandDecode {
		err := cmd.Decode(args)
		if err != nil {
			// check if error contains "authentication failed"
			if strings.Contains(err.Error(), "authentication failed") {
				log.Fatal("Incorrect password")
			} else {
				log.Fatalf("Error decoding image: %v", err)
			}
		}
	} else if command == CommandFit {
		if len(args) == 0 {
			utils.PrintRed("Need image filename")
			return
		}
		err := cmd.Fit(args)
		if err != nil {
			utils.PrintRed(fmt.Sprintf("Error fitting image: %v", err))
		}

	} else {
		utils.ExitWithManual()
	}
}
