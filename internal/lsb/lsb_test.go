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
	origFilename := "../../testdata/original.png"
	dataFilename := "../../testdata/demo.txt"
	outFilename := "../../testdata/test_out.png"

	imgFile, err := os.Open(origFilename)
	assert.NoError(t, err, "Error opening image file")
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	assert.NoError(t, err, "Error decoding image file")

	rgba, ok := img.(*image.RGBA)
	if !ok {
		t.Errorf("Image is not RGBA format")
	}

	dataToEncode, err := os.ReadFile(dataFilename)
	assert.NoError(t, err, "Error reading ", dataFilename)

	err = Encode(rgba, dataToEncode)
	assert.NoError(t, err, "Error encoding data")

	outFile, err := os.Create(outFilename)
	assert.NoError(t, err, "Error creating output file")
	defer outFile.Close()

	err = png.Encode(outFile, rgba)
	assert.NoError(t, err, "Error encoding output image file")

	outFileReRead, err := os.Open(outFilename)
	assert.NoError(t, err, "Error opening output file")
	defer outFileReRead.Close()

	imgDecoded, err := png.Decode(outFileReRead)
	assert.NoError(t, err, "Error decoding output image file")

	rgbaDecoded, ok := imgDecoded.(*image.RGBA)
	if !ok {
		t.Errorf("Decoded image is not RGBA format")
	}

	decodedData := Decode(rgbaDecoded)

	if !bytes.Equal(dataToEncode, decodedData) {
		t.Errorf("Decoded data does not match original data")
	} else {
		fmt.Println("Decoded data matches original data!")
	}

	// cleanup
	err = os.Remove(outFilename)
	assert.NoError(t, err, "Error removing ", outFilename)
}
