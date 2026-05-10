package mapx

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	t.Run("normal map", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		got := Keys(m)
		// 排序以保证断言的稳定性，因为 map 的遍历是无序的
		sort.Strings(got)
		assert.Equal(t, []string{"a", "b", "c"}, got)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		got := Keys(m)
		assert.Equal(t, []string{}, got)
	})

	t.Run("nil map", func(t *testing.T) {
		var m map[string]int
		got := Keys(m)
		assert.Equal(t, []string{}, got)
	})
}

func TestValues(t *testing.T) {
	t.Run("normal map", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		got := Values(m)
		// 排序以保证断言的稳定性
		sort.Ints(got)
		assert.Equal(t, []int{1, 2, 3}, got)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		got := Values(m)
		assert.Equal(t, []int{}, got)
	})

	t.Run("nil map", func(t *testing.T) {
		var m map[string]int
		got := Values(m)
		assert.Equal(t, []int{}, got)
	})
}

func TestMerge(t *testing.T) {
	t.Run("multiple maps", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"b": 3, "c": 4} // "b" 将会覆盖
		m3 := map[string]int{"d": 5}
		got := Merge(m1, m2, m3)
		want := map[string]int{"a": 1, "b": 3, "c": 4, "d": 5}
		assert.Equal(t, want, got)
	})

	t.Run("with empty maps", func(t *testing.T) {
		m1 := map[string]int{"a": 1}
		got := Merge(m1, nil, map[string]int{})
		want := map[string]int{"a": 1}
		assert.Equal(t, want, got)
	})

	t.Run("no arguments", func(t *testing.T) {
		got := Merge[string, int]()
		want := map[string]int{}
		assert.Equal(t, want, got)
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter by value", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		got := Filter(m, func(k string, v int) bool { return v > 1 })
		want := map[string]int{"b": 2, "c": 3}
		assert.Equal(t, want, got)
	})

	t.Run("filter by key", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		got := Filter(m, func(k string, v int) bool { return k == "a" })
		want := map[string]int{"a": 1}
		assert.Equal(t, want, got)
	})

	t.Run("filter all out", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		got := Filter(m, func(k string, v int) bool { return false })
		want := map[string]int{}
		assert.Equal(t, want, got)
	})

	t.Run("empty map", func(t *testing.T) {
		var m map[string]int
		got := Filter(m, func(k string, v int) bool { return true })
		want := map[string]int{}
		assert.Equal(t, want, got)
	})
}

func TestMapValues(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		got := MapValues(m, func(k string, v int) string { return strconv.Itoa(v) })
		want := map[string]string{"a": "1", "b": "2"}
		assert.Equal(t, want, got)
	})

	t.Run("modify value with key", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		got := MapValues(m, func(k string, v int) string { return fmt.Sprintf("%s-%d", k, v) })
		want := map[string]string{"a": "a-1", "b": "b-2"}
		assert.Equal(t, want, got)
	})

	t.Run("empty map", func(t *testing.T) {
		var m map[string]int
		got := MapValues(m, func(k string, v int) string { return strconv.Itoa(v) })
		want := map[string]string{}
		assert.Equal(t, want, got)
	})
}

func TestMapKeys(t *testing.T) {
	t.Run("string to string with prefix", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		got := MapKeys(m, func(k string, v int) string { return "key_" + k })
		want := map[string]int{"key_a": 1, "key_b": 2}
		assert.Equal(t, want, got)
	})

	t.Run("key conflict override", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		// 所有的 key 都被映射为同一个键，最后的会覆盖前面的，但顺序不确定，所以这里只断言长度和其中一个可能的值
		got := MapKeys(m, func(k string, v int) string { return "same_key" })
		assert.Len(t, got, 1)
		// 由于 map 遍历无序，"same_key" 的值可能是 1 也可能是 2
		assert.Contains(t, []int{1, 2}, got["same_key"])
	})

	t.Run("empty map", func(t *testing.T) {
		var m map[string]int
		got := MapKeys(m, func(k string, v int) string { return k + "_new" })
		want := map[string]int{}
		assert.Equal(t, want, got)
	})
}

func TestHas(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}

	t.Run("key exists", func(t *testing.T) {
		assert.True(t, Has(m, "a"))
	})

	t.Run("key not exists", func(t *testing.T) {
		assert.False(t, Has(m, "c"))
	})

	t.Run("nil map", func(t *testing.T) {
		var nilMap map[string]int
		assert.False(t, Has(nilMap, "a"))
	})
}

func TestPick(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	t.Run("pick some keys", func(t *testing.T) {
		got := Pick(m, []string{"a", "c"})
		want := map[string]int{"a": 1, "c": 3}
		assert.Equal(t, want, got)
	})

	t.Run("pick non-existent keys", func(t *testing.T) {
		got := Pick(m, []string{"a", "z"})
		want := map[string]int{"a": 1}
		assert.Equal(t, want, got)
	})

	t.Run("pick empty keys", func(t *testing.T) {
		got := Pick(m, []string{})
		want := map[string]int{}
		assert.Equal(t, want, got)
	})

	t.Run("pick from nil map", func(t *testing.T) {
		var nilMap map[string]int
		got := Pick(nilMap, []string{"a"})
		want := map[string]int{}
		assert.Equal(t, want, got)
	})
}

func TestOmit(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	t.Run("omit some keys", func(t *testing.T) {
		got := Omit(m, []string{"b", "d"})
		want := map[string]int{"a": 1, "c": 3}
		assert.Equal(t, want, got)
	})

	t.Run("omit non-existent keys", func(t *testing.T) {
		got := Omit(m, []string{"b", "z"})
		want := map[string]int{"a": 1, "c": 3, "d": 4}
		assert.Equal(t, want, got)
	})

	t.Run("omit empty keys", func(t *testing.T) {
		got := Omit(m, []string{})
		want := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		assert.Equal(t, want, got)
	})

	t.Run("omit from nil map", func(t *testing.T) {
		var nilMap map[string]int
		got := Omit(nilMap, []string{"a"})
		want := map[string]int{}
		assert.Equal(t, want, got)
	})
}

func TestClone(t *testing.T) {
	t.Run("normal map", func(t *testing.T) {
		original := map[string]int{"a": 1, "b": 2}
		got := Clone(original)
		assert.Equal(t, original, got)
		
		// 验证是浅拷贝且分配了新内存
		got["a"] = 99
		assert.Equal(t, 1, original["a"], "Clone should create a new instance")
	})

	t.Run("empty map", func(t *testing.T) {
		original := map[string]int{}
		got := Clone(original)
		assert.Equal(t, original, got)
		assert.NotNil(t, got)
	})

	t.Run("nil map", func(t *testing.T) {
		var original map[string]int
		got := Clone(original)
		assert.Nil(t, got)
	})
}

func TestInvert(t *testing.T) {
	t.Run("normal invert without conflict", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		got := Invert(m)
		want := map[int]string{1: "a", 2: "b"}
		assert.Equal(t, want, got)
	})

	t.Run("invert with value conflict", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 1}
		got := Invert(m)
		// 由于存在相同的值 1，反转后会被覆盖，长度一定为 1
		assert.Len(t, got, 1)
		// 保留下来的键要么是 "a" 要么是 "b"
		assert.Contains(t, []string{"a", "b"}, got[1])
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		got := Invert(m)
		want := map[int]string{}
		assert.Equal(t, want, got)
	})

	t.Run("nil map", func(t *testing.T) {
		var m map[string]int
		got := Invert(m)
		want := map[int]string{}
		assert.Equal(t, want, got)
	})
}