package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChoose(t *testing.T) {
	t.Run("bool condition is true", func(t *testing.T) {
		got := Choose(true, "yes", "no")
		assert.Equal(t, "yes", got)
	})

	t.Run("bool condition is false", func(t *testing.T) {
		got := Choose(false, 100, 200)
		assert.Equal(t, 200, got)
	})

	t.Run("struct types", func(t *testing.T) {
		type user struct{ name string }
		u1 := user{name: "alice"}
		u2 := user{name: "bob"}
		
		gotTrue := Choose(true, u1, u2)
		assert.Equal(t, u1, gotTrue)

		gotFalse := Choose(false, u1, u2)
		assert.Equal(t, u2, gotFalse)
	})
}

func TestDeepClone(t *testing.T) {
	type Config struct {
		Host string
		Port int
		Tags []string
	}

	t.Run("successful deep clone", func(t *testing.T) {
		src := Config{
			Host: "localhost",
			Port: 8080,
			Tags: []string{"web", "api"},
		}
		var dst Config

		err := DeepClone(src, &dst)
		assert.NoError(t, err)
		assert.Equal(t, src, dst)

		// Verify it's truly deep copy (modifying dst's slice shouldn't affect src)
		dst.Tags[0] = "database"
		assert.Equal(t, "web", src.Tags[0])
		assert.Equal(t, "database", dst.Tags[0])
	})

	t.Run("unmarshal error (dst is not pointer)", func(t *testing.T) {
		src := Config{Host: "localhost"}
		var dst Config

		// dst is passed by value, json.Unmarshal will fail
		err := DeepClone(src, dst)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "反序列化目标对象失败")
	})

	t.Run("marshal error (src cannot be marshaled)", func(t *testing.T) {
		// Functions cannot be JSON marshaled
		src := map[string]interface{}{
			"func": func() {},
		}
		var dst map[string]interface{}

		err := DeepClone(src, &dst)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "序列化源对象失败")
	})
}
