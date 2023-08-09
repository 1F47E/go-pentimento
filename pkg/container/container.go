package container

import (
	"fmt"
	"os"
)

type Crypter interface {
	Encrypt(data []byte, password string) ([]byte, error)
	Decrypt(cipher []byte, password string) ([]byte, error)
}

type Container struct {
	crypter Crypter
}

func New(c Crypter) *Container {
	return &Container{
		crypter: c,
	}
}

// on encoding - reading secret data from file
func (d *Container) EncryptFromFile(filename, password string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading data from file %s: %w", filename, err)
	}
	fmt.Printf("data size:\nplain: %d KB\nplain: %d\n", len(data)/1024, len(data))

	if password != "" {
		fmt.Printf("encrypting with pass: %d\n", len(password))
		cipher, err := d.crypter.Encrypt(data, password)
		if err != nil {
			return nil, fmt.Errorf("error encrypting data: %w", err)
		}
		data = cipher
		fmt.Printf("cipher size: %d KB\n", len(cipher)/1024)
		fmt.Printf("cipher size: %d\n", len(cipher))
	} else {
		fmt.Println("no encryption")
	}
	return data, nil
}

// on decoding - reading secret data from image
func (d *Container) DecryptFromBytes(data []byte, password string) ([]byte, error) {
	fmt.Printf("data size:\ncipher: %d KB\ncipher: %d\n", len(data)/1024, len(data))

	if password != "" {
		fmt.Printf("decrypt with pass: %s\n%X\n", password, password)
		plain, err := d.crypter.Decrypt(data, password)
		if err != nil {
			return nil, fmt.Errorf("error decrypting data: %w", err)
		}
		data = plain
		fmt.Printf("plain size: %d KB\n%d", len(plain)/1024, len(plain))
	} else {
		fmt.Println("no encryption")
	}
	return data, nil
}

func (d *Container) Save(data []byte, filename string) error {
	// TODO: get filename from metadata
	return os.WriteFile(filename, data, 0644)
}
