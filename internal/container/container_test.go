package container

import (
	"errors"
	"testing"
)

// Mock aes crypter
type MockCrypter struct {
	failEncrypt bool
	failDecrypt bool
}

func (m *MockCrypter) Encrypt(data []byte, password string) ([]byte, error) {
	if m.failEncrypt {
		return nil, errors.New("failed to encrypt")
	}
	return append([]byte("enc_"), data...), nil
}

func (m *MockCrypter) Decrypt(cipher []byte, password string) ([]byte, error) {
	if m.failDecrypt {
		return nil, errors.New("failed to decrypt")
	}
	return cipher[4:], nil
}

func TestContainer_EncryptFromFile(t *testing.T) {
	tests := []struct {
		name      string
		file      string
		password  string
		wantError bool
	}{
		{name: "Encrypt demo.txt", file: "../../testdata/demo.txt", password: "pass", wantError: false},
		{name: "Encrypt demo.txt without password", file: "../../testdata/demo.txt", password: "", wantError: false},
		{name: "Encrypt short.txt", file: "../../testdata/short.txt", password: "pass123123123", wantError: false},
		{name: "Encrypt short.txt no pass", file: "../../testdata/short.txt", password: "", wantError: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(&MockCrypter{})
			data, err := c.EncryptFromFile(tt.file, tt.password)
			if (err != nil) != tt.wantError {
				t.Errorf("Error = %v, wantError %v", err, tt.wantError)
				return
			}

			// Check if the data has "enc_" prefix when encrypted
			if tt.password != "" && string(data[:4]) != "enc_" {
				t.Error("Data wasn't encrypted properly")
			}
		})
	}
}
