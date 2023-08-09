package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/1F47E/go-pentimento/cmd"
)

type Command string

const (
	CommandEncode Command = "encode"
	CommandDecode Command = "decode"
	CommandFit    Command = "fit"
)

// func NewData(args []string) (*Data, error) {
// 	if len(args) < 2 {
// 		return nil, fmt.Errorf(usage)
// 	}
// 	imageFilename := args[0]
// 	dataFilename := args[1]

// 	password := ""
// 	if len(os.Args) > 4 {
// 		password = os.Args[4]
// 	}

// 	return &Data{
// 		imageFilename: imageFilename,
// 		dataFilename:  dataFilename,
// 		password:      password,
// 		image:         nil,
// 		data:          nil,
// 	}, nil
// }

// func NewDataFit(args []string) (*Data, error) {
// 	if len(args) == 0 {
// 		return nil, fmt.Errorf(usage)
// 	}
// 	imageFilename := args[0]
// 	return &Data{
// 		imageFilename: imageFilename,
// 		dataFilename:  "",
// 		password:      "",
// 		image:         nil,
// 		data:          nil,
// 	}, nil
// }

// func (d *Data) Encode() {
// 	err := d.openImage()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Encode")

// 	err = d.openData()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// change in place
// 	err = lsb.Encode(d.image, d.data)
// 	if err != nil {
// 		log.Fatalf("Error encoding image: %v", err)
// 	}

// 	outFilename := "hidden.png"
// 	err = saveImage(d.image, outFilename)
// 	if err != nil {
// 		log.Fatalf("Error saving image: %v", err)
// 	}
// 	fmt.Println("Saved image to", outFilename)
// 	os.Exit(0)

// }

// func (d *Data) Decode() {
// 	fmt.Println("Decode")
// 	err := d.openImage()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Encode")

// 	err = d.openData()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	data := lsb.Decode(d.image)
// 	if data == nil {
// 		log.Fatal("No hidden data found")
// 	}
// 	fmt.Printf("Hidden data size: %d bytes\n", len(data))
// 	// save results
// 	resFilename := "decoded.txt"
// 	err = os.WriteFile(resFilename, data, 0644)
// 	if err != nil {
// 		log.Fatalf("Error writing decoded data to file: %v", err)
// 	}
// 	fmt.Println("Data saved to", resFilename)
// }

const usage = `Usage: go run main.go encode|decode|fit <image file> <data file> <password?>`

func main() {
	// var err error
	// debug_createWhiteImage()

	// fmt.Printf("god %d args, got: %v\n", len(os.Args), os.Args)

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(0)
	}

	// get command

	command := Command(os.Args[1])
	// if cmd == CommandFit {
	// 	data, err = NewDataFit(os.Args[2:])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// } else {
	// 	// data, err = NewData(os.Args[2:])
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	// }

	args := os.Args[2:]
	if command == CommandEncode {
		err := cmd.Encode(args)
		if err != nil {
			log.Fatalf("Error encoding image: %v", err)
		}
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
		err := cmd.Fit(args)
		if err != nil {
			log.Fatalf("Error fitting image: %v", err)
		}

	} else {
		fmt.Println(usage)
	}
	os.Exit(0)

}
