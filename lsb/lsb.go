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
	"fmt"
	"image"
	"image/color"
)

var eofMarker byte = 0xFF

func Encode(img *image.RGBA, data []byte) error {
	data = append(data, eofMarker)

	maxLen := MaxBits(img)
	if len(data)*8 > maxLen {
		return fmt.Errorf("data is too large to fit in image")
	}
	perc := float64(len(data)*8) / float64(maxLen) * 100
	fmt.Printf("Encoding %d/%d bytes - %.2f%%\n", len(data), maxLen, perc)

	bitIndex := 0
	byteIndex := 0

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			// current pixel
			c := img.RGBAAt(x, y)
			r, g, b := c.R, c.G, c.B
			// fmt.Printf("Pixel %d,%d: %08b %08b %08b\n", x, y, r, g, b)

			// encode 3 bits to every color channel
			for i := 0; i < 3; i++ {
				// data is exhausted
				if byteIndex >= len(data) {
					return nil
				}

				// ===== get bit from our data byte

				bitPos := uint(bitIndex % 8)
				// right-shifts the byte by bitPos
				// example 01100110 becomes 00011001
				// 0x1 = 00000001 = mask
				bit := (data[byteIndex] >> bitPos) & 00000001

				// 11111110 = 0xFE = 254
				switch i {
				case 0:
					r = (r & 254) | bit
				case 1:
					g = (g & 254) | bit
				case 2:
					b = (b & 254) | bit
				}

				bitIndex++

				if bitIndex%8 == 0 {
					byteIndex++
				}
			}

			// update pixel
			img.SetRGBA(x, y, color.RGBA{r, g, b, c.A})
		}
	}
	return nil
}

func Decode(img *image.RGBA) []byte {
	data := make([]byte, 0)
	var currentByte byte
	bitIndex := 0

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			// current pixel
			c := img.RGBAAt(x, y)
			r, g, b := c.R, c.G, c.B
			// fmt.Printf("Pixel %d,%d: %08b %08b %08b\n", x, y, r, g, b)

			// red, green, blue = 3 bits
			// 0x1 = 00000001 = mask
			for i := 0; i < 3; i++ {
				var bit uint8
				switch i {
				case 0:
					bit = r & 00000001
				case 1:
					bit = g & 00000001
				case 2:
					bit = b & 00000001
				}

				// ====== format out byte with the bit extracted

				// position in the byte where we want to put the bit
				bitPos := uint(bitIndex % 8)
				// bitwise OR operation with the current byte
				// and shift the bit to the correct position
				currentByte |= bit << bitPos

				bitIndex++

				// check for EOR
				if bitIndex%8 == 0 {
					if currentByte == eofMarker {
						return data
					}

					data = append(data, currentByte)
					currentByte = 0
				}
			}
		}
	}
	return data
}

func MaxBits(img *image.RGBA) int {
	return img.Bounds().Dx() * img.Bounds().Dy() * 3
}
