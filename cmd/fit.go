package cmd

import (
	"fmt"

	"github.com/1F47E/go-pentimento/pkg/img"
	"github.com/1F47E/go-pentimento/pkg/lsb"
)

func Fit(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("need image file")
	}
	fImg := args[0]
	img, err := img.Open(fImg)
	if err != nil {
		return fmt.Errorf("error opening image: %w", err)
	}
	bits := lsb.MaxBits(img)
	fmt.Print("Max data size that can be hidden:")
	if bits < 1024 {
		fmt.Printf(" %d bits\n", bits)
	} else if bits < 1024*8 {
		fmt.Printf(" %d Bytes\n", bits/8)
	} else {
		fmt.Printf(" %d KB\n", bits/8/1024)
	}

	return nil
}
