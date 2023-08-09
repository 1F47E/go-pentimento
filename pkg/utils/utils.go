package utils

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func CreateWhiteImage() {
	// create white image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)

	// save image
	f, err := os.Create("white.png")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer f.Close()

	// encode image
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("error encoding image: %v", err)
	}
}
