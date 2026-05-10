package hashx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	t.Run("normal string", func(t *testing.T) {
		input := "hello world"
		// MD5 of "hello world" is 5eb63bbbe01eeed093cb22bb8f5acdc3
		expected := "5eb63bbbe01eeed093cb22bb8f5acdc3"
		got := MD5(input)
		assert.Equal(t, expected, got)
	})

	t.Run("empty string", func(t *testing.T) {
		input := ""
		// MD5 of empty string is d41d8cd98f00b204e9800998ecf8427e
		expected := "d41d8cd98f00b204e9800998ecf8427e"
		got := MD5(input)
		assert.Equal(t, expected, got)
	})

	t.Run("special characters", func(t *testing.T) {
		input := "!@#$%^&*()_+{}|:\"<>?~`-=[]\\;',./"
		// MD5 should handle any characters consistently
		got1 := MD5(input)
		got2 := MD5(input)
		assert.NotEmpty(t, got1)
		assert.Equal(t, got1, got2) // MD5 is deterministic
	})

	t.Run("chinese characters", func(t *testing.T) {
		input := "你好，世界"
		got1 := MD5(input)
		got2 := MD5(input)
		assert.NotEmpty(t, got1)
		assert.Equal(t, got1, got2) // MD5 is deterministic
	})

	t.Run("long string", func(t *testing.T) {
		input := string(make([]byte, 1024*1024)) // 1MB of null bytes
		got := MD5(input)
		assert.NotEmpty(t, got)
		assert.Len(t, got, 32) // MD5 hex string is always 32 characters long
	})
}
