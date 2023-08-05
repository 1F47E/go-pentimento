package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func maxHiddenBits(img *image.RGBA) int {
	return img.Bounds().Dx() * img.Bounds().Dy() * 3
}

func main() {
	// get filename from arg
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]

	// open image
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// TODO: work with any image not just PNG

	// read image
	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	// convert to RGBA
	// convert to RGBA
	bounds := img.Bounds()
	imgRGBA := image.NewRGBA(bounds)
	draw.Draw(imgRGBA, bounds, img, bounds.Min, draw.Src)

	// calc size of data that can be hidden
	maxBits := maxHiddenBits(imgRGBA)
	fmt.Printf("Max data size that can be hidden: %dB, %d Bytes, %d KB\n", maxBits, maxBits/8, maxBits/8/1024)

}
