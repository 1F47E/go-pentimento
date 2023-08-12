package img

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func Open(filename string) (*image.RGBA, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening image: %v", err)
	}
	defer f.Close()

	// detect image type and decode
	img := image.Image(nil)
	ext := strings.ToLower(filepath.Ext(filename))
	if ext == ".png" {
		img, err = png.Decode(f)
	} else if ext == ".jpg" || ext == ".jpeg" {
		img, err = jpeg.Decode(f)
	} else if ext == ".gif" {
		img, err = gif.Decode(f)
	} else {
		return nil, fmt.Errorf("unsupported image type: %s", ext)
	}

	if err != nil {
		return nil, fmt.Errorf("error decoding image: %v", err)
	}

	// convert to RGBA and return pointer
	bounds := img.Bounds()
	imgRGBA := image.NewRGBA(bounds)
	draw.Draw(imgRGBA, bounds, img, bounds.Min, draw.Src)
	return imgRGBA, nil
}

func Save(img *image.RGBA, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return fmt.Errorf("error encoding image: %v", err)
	}
	return nil
}
