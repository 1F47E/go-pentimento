//
// encoding 3 bits per pixel
//
// red channel - first bit
// green channel - second bit
// blue channel - third bit

// Example:
//
// R = 11010101
// G = 10101110
// B = 11110011
//
// data = 10110011
//
// in LSB we take the least significant bit from each color value
// and replace it with the bit from our data (right one)
//
// res after encoding the first three bits of our data
// R = 11010100 1 -> 0
// G = 10101111 0 -> 1
// B = 11110010 1 -> 0
//

package lsb

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"image"
	"image/color"
)

var eofMarker = []byte{0xFF, 0xFF, 0xFF}

func Encode(img *image.RGBA, data []byte) error {
	data = append(data, eofMarker...)

	// check if data fits in image
	maxLen := MaxBits(img)
	if len(data)*8 > maxLen {
		return fmt.Errorf("data is too large to fit in image")
	}
	perc := float64(len(data)*8) / float64(maxLen) * 100
	fmt.Printf("Encoding %d/%d bytes - %.2f%%\n", len(data), maxLen, perc)

	// track the index of corrent write bit
	bitIndex := 0

	// iterace over pixels
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			// current pixel
			c := img.RGBAAt(x, y)
			r, g, b := c.R, c.G, c.B

			// decide what pixl bit to use, classic LSB or shift right by 1
			// its not exactly 2LSB, because we still using only 1 bit, not 2
			// to hide LSB from detection
			use2LSB := useSecondBit(x, y)
			var mask uint8
			if use2LSB {
				mask = 0b11111101
			} else {
				mask = 0b11111110
			}

			// iterate over RGB channels
			for i := 0; i < 3; i++ {
				if bitIndex >= len(data)*8 {
					// since we writing to the image in place - just quit
					return nil
				}

				// get the bit from data
				byteIndex := bitIndex / 8
				bitPos := bitIndex % 8

				// Get the bit value
				bit := (data[byteIndex] >> bitPos) & 1
				if use2LSB {
					bit = bit << 1
				}

				bitIndex++

				switch i {
				case 0:
					// fmt.Printf("-red before masking: %08b\n", r)

					// 1 - zero out the bit we want to change
					// original state of r (for example):    0b10010110
					// mask:                               & 0b11111101
					// result after masking:                 0b10010100
					r = r & mask

					// 2 - apply the bit to the channel
					// current state of r (after masking):  0b10010100
					// bit to set for example 1           | 0b00000010
					// res after ORing the bit:             0b10010110
					r |= bit

					// fmt.Printf("+red after masking: %08b\n", r)

				case 1:
					g = (g & mask) | bit
				case 2:
					b = (b & mask) | bit
				}
			}

			// apply pixel
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	if bitIndex < len(data)*8 {
		return fmt.Errorf("not all data was encoded")
	}

	return nil
}

func Decode(img *image.RGBA) []byte {
	data := []byte{}
	bitIndex := 0

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			c := img.RGBAAt(x, y)
			r, g, b := c.R, c.G, c.B

			use2LSB := useSecondBit(x, y)

			for i := 0; i < 3; i++ {
				var channelValue byte
				switch i {
				case 0:
					channelValue = r
				case 1:
					channelValue = g
				case 2:
					channelValue = b
				}

				// for 2LSB shift bit right
				if use2LSB {
					channelValue = channelValue >> 1
				}

				// get our bit
				bit := channelValue & 0b00000001

				// prepare
				byteIndex := bitIndex / 8
				if byteIndex >= len(data) {
					data = append(data, 0)
				}

				// apply bit
				// 00010011 (data[byteIndex] before)
				// 00001000 (bit shifted left (bitIndex % 8) positions)
				// --------- (Bitwise OR operation)
				// 00011011 (data[byteIndex] after)
				data[byteIndex] |= bit << (bitIndex % 8)

				bitIndex++
			}
			// fmt.Printf("x:%d y:%d %08b %08b %08b\n", x, y, r, g, b)
		}
	}

	// cut off on EOF
	eofIndex := bytes.Index(data, eofMarker)
	if eofIndex != -1 {
		data = data[:eofIndex]
	}

	return data
}

func MaxBits(img *image.RGBA) int {
	return img.Bounds().Dx() * img.Bounds().Dy() * 3
}

// Hiding LSB from detection via custom algorithm
// via switching between 1 and 2 bit pos. last (LSB) and second to last (2LSB)
// based on a hash from pixel pos
func hashPos(x, y int) int {
	pix := fmt.Sprintf("%d", x+y)
	hash := fnv.New64a()
	hash.Write([]byte(pix))
	return int(hash.Sum64())
}

// decide which LSB to use
func useSecondBit(x, y int) bool {
	h := hashPos(x, y)
	pos := h % 2
	return pos == 1
}
