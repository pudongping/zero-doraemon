package hashx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	t.Run("normal password", func(t *testing.T) {
		password := "my_secret_password_123!"
		hash, err := HashPassword(password)
		
		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
		assert.NotEqual(t, password, hash) // hash must be different from plain text
	})

	t.Run("empty password", func(t *testing.T) {
		password := ""
		hash, err := HashPassword(password)
		
		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
	})

	t.Run("very long password", func(t *testing.T) {
		// bcrypt max length is 72 bytes
		// If password is longer than 72 bytes, bcrypt truncates it or returns an error depending on the implementation
		password := string(make([]byte, 100))
		hash, err := HashPassword(password)
		
		// The golang.org/x/crypto/bcrypt implementation will truncate passwords > 72 bytes but still return no error
		if err == nil {
			assert.NotEmpty(t, hash)
		} else {
			assert.ErrorIs(t, err, bcrypt.ErrPasswordTooLong)
		}
	})

	t.Run("different hashes for same password", func(t *testing.T) {
		// bcrypt uses a random salt, so hashing the same password twice should yield different results
		password := "password"
		hash1, err1 := HashPassword(password)
		hash2, err2 := HashPassword(password)
		
		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.NotEqual(t, hash1, hash2)
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Run("correct password", func(t *testing.T) {
		password := "my_secret_password"
		hash, err := HashPassword(password)
		assert.NoError(t, err)

		isMatch := CheckPasswordHash(hash, password)
		assert.True(t, isMatch)
	})

	t.Run("incorrect password", func(t *testing.T) {
		password := "my_secret_password"
		hash, err := HashPassword(password)
		assert.NoError(t, err)

		isMatch := CheckPasswordHash(hash, "wrong_password")
		assert.False(t, isMatch)
	})

	t.Run("empty password match", func(t *testing.T) {
		password := ""
		hash, err := HashPassword(password)
		assert.NoError(t, err)

		isMatch := CheckPasswordHash(hash, password)
		assert.True(t, isMatch)
	})

	t.Run("invalid hash format", func(t *testing.T) {
		isMatch := CheckPasswordHash("not_a_valid_bcrypt_hash", "password")
		assert.False(t, isMatch)
	})
}
