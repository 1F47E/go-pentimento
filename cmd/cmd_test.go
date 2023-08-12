package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testdata        = "../testdata"
	results         = "decoded-data.txt"
	secretText      = "demo.txt"
	secretTextShort = "short.txt"

	imgPngOriginal = "original.png"
	imgJpgOriginal = "original.jpg"
	imgGifOriginal = "original.gif"
	imgPngEncoded  = "encoded-original.png"
	imgJpgEncoded  = "encoded-original.png"
	imgGifEncoded  = "encoded-original.png"

	// decoding only dest
	defaultPassword        = "asdfasdfasdf"
	imgEncoded             = "encoded-original.png"
	imgEncodedNoPassword   = "encoded-demo.txt-nopassword.png"
	imgEncodedWithPassword = "encoded-demo.txt-password-asdfasdfasdf.png"
	img404                 = "nonexistent.png"
)

func randomPassword() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	l := r.Intn(200) + 1
	s := make([]rune, l)
	for i := range s {
		s[i] = letters[r.Intn(len(letters))]
	}
	return string(s)
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantError bool
	}{
		{
			name:      "Not Enough Args",
			args:      []string{filepath.Join(testdata, imgPngOriginal)},
			wantError: true,
		},
		{
			name:      "Non-existent Image",
			args:      []string{filepath.Join(testdata, img404), filepath.Join(testdata, secretText)},
			wantError: true,
		},
		{
			name:      "Valid PNG Image And Data Without Password",
			args:      []string{filepath.Join(testdata, imgPngOriginal), filepath.Join(testdata, secretText)},
			wantError: false,
		},
		{
			name:      "Valid JPG Image And Data Without Password",
			args:      []string{filepath.Join(testdata, imgJpgOriginal), filepath.Join(testdata, secretText)},
			wantError: false,
		},
		{
			name:      "Valid GIF Image And Data Without Password",
			args:      []string{filepath.Join(testdata, imgGifOriginal), filepath.Join(testdata, secretText)},
			wantError: false,
		},
		{
			name:      "Valid Image And Data With Password",
			args:      []string{filepath.Join(testdata, imgPngOriginal), filepath.Join(testdata, secretText), defaultPassword},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Encode(tt.args)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				origFilename := tt.args[0]
				encodedFile := fmt.Sprintf("encoded-%s.png", filepath.Base(origFilename[:len(origFilename)-len(filepath.Ext(origFilename))]))
				assert.FileExists(t, encodedFile, "The encoded image does not exist at the expected location")
			}
		})
	}

	// cleanup
	err := os.Remove(imgEncoded)
	assert.NoError(t, err, "Error removing encoded-original.png")
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantError   bool
		checkResult bool
	}{
		{
			name:        "No Args",
			args:        []string{},
			wantError:   true,
			checkResult: false,
		},

		// decode previously encoded img without password
		{
			name:        "Only Image Path",
			args:        []string{filepath.Join(testdata, imgEncodedNoPassword)},
			wantError:   false,
			checkResult: true,
		},

		// decode previously encoded img with password asdfasdfasdf
		// content is testdata/demo.txt
		{
			name:        "Image Path With Password",
			args:        []string{filepath.Join(testdata, imgEncodedWithPassword), defaultPassword},
			wantError:   false,
			checkResult: true,
		},
		{
			name:        "Non-existent Image",
			args:        []string{filepath.Join(testdata, img404)},
			wantError:   true,
			checkResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Decode(tt.args)
			if tt.wantError {
				assert.Error(t, err, "expected an error but got none")
			} else if err != nil {
				t.Fatalf("didn't expect an error but got %s", err.Error())
			}

			// check if decoded data matches original data
			if tt.checkResult {
				// read original secret data
				secretData, err := os.ReadFile(filepath.Join(testdata, secretText))
				assert.NoError(t, err, "Error reading ", secretText)

				// read decoded data from file
				decodedData, err := os.ReadFile(results)
				assert.NoError(t, err, "Error reading ", results)

				// compare
				assert.True(t, string(secretData) == string(decodedData), "Original and decoded do not match: %s != %s", secretData, decodedData)
			}

		})
	}

	// cleanup
	err := os.Remove(results)
	if err != nil {
		t.Fatalf("error removing %s: %s", results, err.Error())
	}
}

func TestEncodeDecodeWithoutPassword(t *testing.T) {
	// secretText := secretTextShort
	// encode without password
	argsEncode := []string{filepath.Join(testdata, imgPngOriginal), filepath.Join(testdata, secretText)}
	err := Encode(argsEncode)
	assert.NoError(t, err, "Error during encoding without password")

	// decode
	argsDecode := []string{imgEncoded}
	err = Decode(argsDecode)
	assert.NoError(t, err, "Error during decoding without password")

	// compare
	originalData, err := os.ReadFile(filepath.Join(testdata, secretText))
	assert.NoError(t, err, "Error reading original secret text")
	decodedData, err := os.ReadFile(results)
	assert.NoError(t, err, "Error reading decoded results")
	assert.Equal(t, originalData, decodedData, "Original and decoded data are not equal for encoding/decoding without password")

	// cleanup
	err = os.Remove(imgEncoded)
	assert.NoError(t, err, "Error removing encoded image")
}

func TestEncodeDecodeWithRandomPassword(t *testing.T) {
	password := randomPassword()
	// password := "123"
	// secretText := secretTextShort

	// encode with random password
	argsEncode := []string{filepath.Join(testdata, imgPngOriginal), filepath.Join(testdata, secretText), password}
	err := Encode(argsEncode)
	assert.NoError(t, err, "Error during encoding with random password")

	// decode with the same password
	// image is in the same dir as test
	argsDecode := []string{imgEncoded, password}
	err = Decode(argsDecode)
	assert.NoError(t, err, "Error during decoding with random password")

	// compare decoded data to original data
	originalData, err := os.ReadFile(filepath.Join(testdata, secretText))
	assert.NoError(t, err, "Error reading original secret text")
	decodedData, err := os.ReadFile(results)
	assert.NoError(t, err, "Error reading decoded results")
	assert.Equal(t, originalData, decodedData, "Original and decoded data are not equal for encoding/decoding with random password")

	// cleanup
	err = os.Remove(imgEncoded)
	assert.NoError(t, err, "Error removing encoded image")
	err = os.Remove(results)
	assert.NoError(t, err, "Error removing decoded results")
}
