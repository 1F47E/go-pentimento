package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/1F47E/go-pentimento/pkg/lsb"
)

type Command string

const (
	Encode Command = "encode"
	Decode Command = "decode"
	Fit    Command = "fit"
)

const usage = `Usage: go run main.go encode|decode|fit <filename> text`

func main() {
	// debug_createWhiteImage()
	// panic("debug")

	// get filename from arg
	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}
	cmd := Command(os.Args[1])
	if cmd != Encode && cmd != Decode && cmd != Fit {
		fmt.Println(usage)
		os.Exit(1)
	}
	filename := os.Args[2]

	// get text from arg only if encoding
	datafile := ""
	if cmd == Encode {
		if len(os.Args) < 4 {
			fmt.Println(usage)
			os.Exit(1)
		}
		datafile = os.Args[3]
	}

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
	bounds := img.Bounds()
	imgRGBA := image.NewRGBA(bounds) // returns pointer
	draw.Draw(imgRGBA, bounds, img, bounds.Min, draw.Src)

	// ENCODE
	if cmd == Encode {

		// read data file
		data, err := os.ReadFile(datafile)
		if err != nil {
			panic(fmt.Errorf("error reading data file: %v", err))
		}
		fmt.Printf("Data to hide: %d bytes\n", len(data))

		// change in place
		err = lsb.Encode(imgRGBA, data)
		if err != nil {
			panic(err)
		}

		outFilename := "out.png"
		err = saveImage(imgRGBA, outFilename)
		if err != nil {
			panic(err)
		}
		fmt.Println("Saved image to", outFilename)
		os.Exit(0)
	}

	// DECODE
	if cmd == Decode {
		data := lsb.Decode(imgRGBA)
		fmt.Printf("Hidden data: %s\n", string(data))
		os.Exit(0)
	}

	// FIT
	if cmd == Fit {
		maxBits := lsb.MaxBits(imgRGBA)
		fmt.Printf("Max data size that can be hidden: %dB, %d Bytes, %d KB\n", maxBits, maxBits/8, maxBits/8/1024)
	}
}

func saveImage(img *image.RGBA, filename string) error {
	// save image
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer f.Close()

	// encode image
	err = png.Encode(f, img)
	if err != nil {
		return fmt.Errorf("error encoding image: %v", err)
	}
	return nil
}

func debugCreateWhiteImage() {
	// create white image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)

	// save image
	f, err := os.Create("white.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// encode image
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
