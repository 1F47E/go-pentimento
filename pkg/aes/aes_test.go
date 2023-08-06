package myaes

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAESCrypter(t *testing.T) {
	t.Run("Encryption and Decryption", func(t *testing.T) {
		var a = new(AEScrypter)

		// TODO: test with loading session from file

		data := []byte("Hello World!")
		pass := "password"

		// Encrypt a message
		cipher, err := a.Encrypt(data, pass)
		assert.NoError(t, err, "Encrypt error")

		// Decrypt the message
		plain, err := a.Decrypt(cipher, pass)
		assert.NoError(t, err, "Decrypt error")

		// Test assert original and decoded
		assert.True(t, bytes.Equal(plain, data), "Original and decoded do not match: %s != %s", plain, data)
	})
}
