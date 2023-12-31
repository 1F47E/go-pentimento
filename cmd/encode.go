package cmd

import (
	"fmt"
	"path/filepath"

	myaes "github.com/1F47E/go-pentimento/internal/aes"
	"github.com/1F47E/go-pentimento/internal/container"
	"github.com/1F47E/go-pentimento/internal/img"
	"github.com/1F47E/go-pentimento/internal/lsb"
)

func Encode(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("need more data")
	}
	fImg := args[0]
	fData := args[1]
	password := ""
	if len(args) > 2 {
		password = string(args[2])
	}

	ig, err := img.Open(fImg)
	if err != nil {
		return fmt.Errorf("error opening image: %w", err)
	}

	crypter := myaes.New()
	box := container.New(crypter)

	data, err := box.EncryptFromFile(fData, password)
	if err != nil {
		return fmt.Errorf("error processing data: %w", err)
	}

	// in place encode to the img
	err = lsb.Encode(ig, data)
	if err != nil {
		return fmt.Errorf("error encoding image: %w", err)
	}

	// get filename without extension
	filename := filepath.Base(fImg)
	ext := filepath.Ext(filename)
	base := filename[:len(filename)-len(ext)]
	err = img.Save(ig, fmt.Sprintf("encoded-%s.png", base))
	if err != nil {
		return fmt.Errorf("error saving image: %w", err)
	}
	return nil
}
