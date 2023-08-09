package cmd

import (
	"fmt"

	myaes "github.com/1F47E/go-pentimento/pkg/aes"
	"github.com/1F47E/go-pentimento/pkg/container"
	"github.com/1F47E/go-pentimento/pkg/img"
	"github.com/1F47E/go-pentimento/pkg/lsb"
)

func Decode(args []string) error {
	var err error

	if len(args) == 0 {
		return fmt.Errorf("need image data")
	}
	fImg := args[0]
	// fData := args[1]
	password := ""
	if len(args) > 1 {
		password = args[1]
	}
	img, err := img.Open(fImg)
	if err != nil {
		return fmt.Errorf("error opening image: %w", err)
	}
	userData := lsb.Decode(img)
	fmt.Printf("hidden data extracted: %d Kb\n", len(userData)/1024)

	// decrypt data if needed
	crypter := myaes.New()
	box := container.New(crypter)
	data, err := box.DecryptFromBytes(userData, password)
	if err != nil {
		return fmt.Errorf("error processing data: %w", err)
	}

	// TODO: get image name from metadata
	err = box.Save(data, "decoded-data.txt")
	if err != nil {
		return fmt.Errorf("error saving data: %w", err)
	}

	return nil
}
