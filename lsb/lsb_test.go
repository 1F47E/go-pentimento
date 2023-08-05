package lsb

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	imgFile, err := os.Open("../original.png")
	assert.NoError(t, err, "Error opening image file")
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	assert.NoError(t, err, "Error decoding image file")

	rgba, ok := img.(*image.RGBA)
	if !ok {
		t.Errorf("Image is not RGBA format")
	}

	textFile, err := os.ReadFile("../README.md")
	assert.NoError(t, err, "Error reading README.md")

	err = Encode(rgba, textFile)
	assert.NoError(t, err, "Error encoding data")

	outFile, err := os.Create("out.png")
	assert.NoError(t, err, "Error creating output file")
	defer outFile.Close()

	err = png.Encode(outFile, rgba)
	assert.NoError(t, err, "Error encoding output image file")

	outFileReRead, err := os.Open("out.png")
	assert.NoError(t, err, "Error opening output file")
	defer outFileReRead.Close()

	imgDecoded, err := png.Decode(outFileReRead)
	assert.NoError(t, err, "Error decoding output image file")

	rgbaDecoded, ok := imgDecoded.(*image.RGBA)
	if !ok {
		t.Errorf("Decoded image is not RGBA format")
	}

	decodedData := Decode(rgbaDecoded)

	if !bytes.Equal(textFile, decodedData) {
		t.Errorf("Decoded data does not match original data")
	} else {
		fmt.Println("Decoded data matches original data!")
	}

	// cleanup
	err = os.Remove("out.png")
	assert.NoError(t, err, "Error removing out.png")
}
