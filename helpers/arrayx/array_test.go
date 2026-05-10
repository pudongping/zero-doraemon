package arrayx

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIn(t *testing.T) {
	tests := []struct {
		name    string
		element int
		slice   []int
		want    bool
	}{
		{"exists", 2, []int{1, 2, 3}, true},
		{"not exists", 4, []int{1, 2, 3}, false},
		{"empty slice", 1, []int{}, false},
		{"nil slice", 1, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := In(tt.element, tt.slice)
			assert.Equal(t, tt.want, got)
		})
	}

	// Test with strings
	got := In("b", []string{"a", "b", "c"})
	assert.True(t, got)
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"empty slice", []int{}, []int{}},
		{"nil slice", nil, []int{}},
		{"no duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"with duplicates", []int{1, 2, 2, 3, 1}, []int{1, 2, 3}},
		{"all duplicates", []int{1, 1, 1}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Unique(tt.slice)
			assert.Equal(t, tt.want, got)
		})
	}

	// Test with strings
	got := Unique([]string{"a", "b", "a"})
	assert.Equal(t, []string{"a", "b"}, got)
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name   string
		slices [][]int
		want   []int
	}{
		{"no slices", nil, []int{}},
		{"empty slices", [][]int{{}, {}}, []int{}},
		{"single slice", [][]int{{1, 2}}, []int{1, 2}},
		{"multiple slices", [][]int{{1, 2}, {3, 4}, {5}}, []int{1, 2, 3, 4, 5}},
		{"with empty slice in middle", [][]int{{1}, {}, {2}}, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.slices...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDiff(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   []int
	}{
		{"slice1 empty", []int{}, []int{1, 2}, []int{}},
		{"slice2 empty", []int{1, 2}, []int{}, []int{1, 2}},
		{"normal diff", []int{1, 2, 3, 4}, []int{2, 4, 5}, []int{1, 3}},
		{"no diff (subset)", []int{2, 4}, []int{1, 2, 3, 4}, []int{}},
		{"all diff", []int{1, 2}, []int{3, 4}, []int{1, 2}},
		{"with duplicates in slice1", []int{1, 2, 2, 3}, []int{2}, []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Diff(tt.slice1, tt.slice2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   []int
	}{
		{"slice1 empty", []int{}, []int{1, 2}, []int{}},
		{"slice2 empty", []int{1, 2}, []int{}, []int{}},
		{"normal intersect", []int{1, 2, 3, 2}, []int{2, 4, 3}, []int{2, 3}},
		{"no intersect", []int{1, 2}, []int{3, 4}, []int{}},
		{"all intersect", []int{1, 2}, []int{1, 2}, []int{1, 2}},
		{"duplicates in slice1", []int{1, 1, 2}, []int{1, 3}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersect(tt.slice1, tt.slice2)
			assert.Equal(t, tt.want, got)
		})
	}
}

type testEmployee struct {
	ID         int
	Name       string
	Department string
	Age        int
	IsAdmin    bool
}

func TestFilter(t *testing.T) {
	t.Run("filter int slice", func(t *testing.T) {
		tests := []struct {
			name      string
			slice     []int
			predicate func(int) bool
			want      []int
		}{
			{"empty slice", []int{}, func(v int) bool { return true }, []int{}},
			{"keep all", []int{1, 2, 3}, func(v int) bool { return true }, []int{1, 2, 3}},
			{"remove all", []int{1, 2, 3}, func(v int) bool { return false }, []int{}},
			{"keep even", []int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 }, []int{2, 4}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Filter(tt.slice, tt.predicate)
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("filter struct slice", func(t *testing.T) {
		employees := []testEmployee{
			{1, "Alice", "HR", 25, false},
			{2, "Bob", "IT", 35, true},
			{3, "Charlie", "IT", 28, false},
		}
		// 获取所有 IT 部门的员工
		gotIT := Filter(employees, func(e testEmployee) bool { return e.Department == "IT" })
		assert.Len(t, gotIT, 2)
		assert.Equal(t, "Bob", gotIT[0].Name)

		// 获取所有管理员
		gotAdmins := Filter(employees, func(e testEmployee) bool { return e.IsAdmin })
		assert.Len(t, gotAdmins, 1)
		assert.Equal(t, "Bob", gotAdmins[0].Name)
	})
}

func TestMap(t *testing.T) {
	// Test int to string
	t.Run("int to string", func(t *testing.T) {
		slice := []int{1, 2, 3}
		want := []string{"1", "2", "3"}
		got := Map(slice, func(v int) string { return strconv.Itoa(v) })
		assert.Equal(t, want, got)
	})

	// Test empty
	t.Run("empty slice", func(t *testing.T) {
		var slice []int
		want := []string{}
		got := Map(slice, func(v int) string { return strconv.Itoa(v) })
		assert.Equal(t, want, got)
	})

	// Test string to int
	t.Run("string to length", func(t *testing.T) {
		slice := []string{"a", "bb", "ccc"}
		want := []int{1, 2, 3}
		got := Map(slice, func(v string) int { return len(v) })
		assert.Equal(t, want, got)
	})

	// Test extract struct field (pluck)
	t.Run("extract struct field", func(t *testing.T) {
		employees := []testEmployee{
			{1, "Alice", "HR", 25, false},
			{2, "Bob", "IT", 35, true},
			{3, "Charlie", "IT", 28, false},
		}
		// 提取所有员工的 ID 组成新切片
		gotIDs := Map(employees, func(e testEmployee) int { return e.ID })
		assert.Equal(t, []int{1, 2, 3}, gotIDs)

		// 提取所有员工的名字组成新切片
		gotNames := Map(employees, func(e testEmployee) string { return e.Name })
		assert.Equal(t, []string{"Alice", "Bob", "Charlie"}, gotNames)
	})
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"even elements", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"odd elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Need to clone because Reverse modifies in-place
			input := Clone(tt.slice)
			got := Reverse(input)
			assert.Equal(t, tt.want, got)
			// check if original slice was modified (in-place)
			if len(tt.slice) > 1 {
				assert.NotEqual(t, tt.slice, input, "Reverse() did not modify slice in-place")
			}
		})
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		size  int
		want  [][]int
	}{
		{"empty slice", []int{}, 2, [][]int{}},
		{"size <= 0", []int{1, 2}, 0, [][]int{}},
		{"size > len", []int{1, 2}, 5, [][]int{{1, 2}}},
		{"exact multiple", []int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{"not exact multiple", []int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.slice, tt.size)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestContainsBy(t *testing.T) {
	t.Run("contains int slice", func(t *testing.T) {
		tests := []struct {
			name      string
			slice     []int
			predicate func(int) bool
			want      bool
		}{
			{"empty slice", []int{}, func(v int) bool { return true }, false},
			{"contains true", []int{1, 2, 3}, func(v int) bool { return v > 2 }, true},
			{"contains false", []int{1, 2, 3}, func(v int) bool { return v > 5 }, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := ContainsBy(tt.slice, tt.predicate)
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("contains struct slice", func(t *testing.T) {
		employees := []testEmployee{
			{1, "Alice", "HR", 25, false},
			{2, "Bob", "IT", 35, true},
		}
		// 检查是否存在管理员
		hasAdmin := ContainsBy(employees, func(e testEmployee) bool { return e.IsAdmin })
		assert.True(t, hasAdmin)

		// 检查是否存在财务部员工
		hasFinance := ContainsBy(employees, func(e testEmployee) bool { return e.Department == "Finance" })
		assert.False(t, hasFinance)
	})
}

func TestGroupBy(t *testing.T) {
	t.Run("group by parity", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		want := map[string][]int{
			"odd":  {1, 3},
			"even": {2, 4},
		}
		got := GroupBy(slice, func(v int) string {
			if v%2 == 0 {
				return "even"
			}
			return "odd"
		})
		assert.Equal(t, want, got)
	})

	t.Run("empty slice", func(t *testing.T) {
		var slice []int
		want := map[string][]int{}
		got := GroupBy(slice, func(v int) string { return "any" })
		assert.Equal(t, want, got)
	})

	t.Run("group structs by field", func(t *testing.T) {
		employees := []testEmployee{
			{1, "Alice", "HR", 25, false},
			{2, "Bob", "IT", 35, true},
			{3, "Charlie", "IT", 28, false},
		}
		// 按照部门进行分组
		got := GroupBy(employees, func(e testEmployee) string { return e.Department })
		assert.Len(t, got, 2)
		assert.Len(t, got["IT"], 2)
		assert.Len(t, got["HR"], 1)
		assert.Equal(t, "Bob", got["IT"][0].Name)
	})
}

type testUser struct {
	ID   int
	Name string
}

func TestKeyBy(t *testing.T) {
	t.Run("normal struct slice", func(t *testing.T) {
		users := []testUser{{1, "Alice"}, {2, "Bob"}}
		want := map[int]testUser{
			1: {1, "Alice"},
			2: {2, "Bob"},
		}
		got := KeyBy(users, func(u testUser) int { return u.ID })
		assert.Equal(t, want, got)
	})

	t.Run("override keys", func(t *testing.T) {
		users := []testUser{{1, "Alice"}, {1, "Alice Updated"}}
		want := map[int]testUser{
			1: {1, "Alice Updated"},
		}
		got := KeyBy(users, func(u testUser) int { return u.ID })
		assert.Equal(t, want, got)
	})

	t.Run("empty slice", func(t *testing.T) {
		var users []testUser
		want := map[int]testUser{}
		got := KeyBy(users, func(u testUser) int { return u.ID })
		assert.Equal(t, want, got)
	})

	t.Run("index structs by string field", func(t *testing.T) {
		employees := []testEmployee{
			{1, "Alice", "HR", 25, false},
			{2, "Bob", "IT", 35, true},
		}
		// 将员工列表转换成以 Name 为键的 map (适用于根据名字快速查询)
		got := KeyBy(employees, func(e testEmployee) string { return e.Name })
		assert.Len(t, got, 2)
		assert.Equal(t, 35, got["Bob"].Age)
		assert.Equal(t, 1, got["Alice"].ID)
	})
}

func TestClone(t *testing.T) {
	t.Run("normal slice", func(t *testing.T) {
		original := []int{1, 2, 3}
		got := Clone(original)
		assert.Equal(t, original, got)
		// Verify it's a new instance by modifying the clone
		got[0] = 99
		assert.NotEqual(t, 99, original[0], "Clone() did not create a new slice instance")
	})

	t.Run("empty slice", func(t *testing.T) {
		original := []int{}
		got := Clone(original)
		assert.Equal(t, original, got)
		assert.NotNil(t, got)
	})

	t.Run("nil slice", func(t *testing.T) {
		var original []int
		got := Clone(original)
		assert.Nil(t, got)
	})
}

func TestMax(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		got := Max([]int{1, 5, 3, 9, 2})
		assert.Equal(t, 9, got)
	})

	t.Run("float slice", func(t *testing.T) {
		got := Max([]float64{1.1, 5.5, 3.3})
		assert.Equal(t, 5.5, got)
	})

	t.Run("negative numbers", func(t *testing.T) {
		got := Max([]int{-5, -1, -3})
		assert.Equal(t, -1, got)
	})

	t.Run("panic on empty", func(t *testing.T) {
		assert.PanicsWithValue(t, "arrayx: Max called on an empty slice", func() {
			Max([]int{})
		})
	})
}

func TestMin(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		got := Min([]int{4, 5, 1, 9, 2})
		assert.Equal(t, 1, got)
	})

	t.Run("float slice", func(t *testing.T) {
		got := Min([]float64{4.4, 5.5, 1.1})
		assert.Equal(t, 1.1, got)
	})

	t.Run("negative numbers", func(t *testing.T) {
		got := Min([]int{-5, -1, -3})
		assert.Equal(t, -5, got)
	})

	t.Run("panic on empty", func(t *testing.T) {
		assert.PanicsWithValue(t, "arrayx: Min called on an empty slice", func() {
			Min([]int{})
		})
	})
}

func TestSum(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4})
		assert.Equal(t, 10, got)
	})

	t.Run("float slice", func(t *testing.T) {
		got := Sum([]float64{1.5, 2.5})
		assert.Equal(t, 4.0, got)
	})

	t.Run("empty slice", func(t *testing.T) {
		got := Sum([]int{})
		assert.Equal(t, 0, got)
	})
}

func TestAvg(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		got := Avg([]int{1, 2, 3, 4})
		assert.Equal(t, 2.5, got)
	})

	t.Run("float slice", func(t *testing.T) {
		got := Avg([]float64{1.5, 2.5, 5.0})
		assert.Equal(t, 3.0, got)
	})

	t.Run("panic on empty", func(t *testing.T) {
		assert.PanicsWithValue(t, "arrayx: Avg called on an empty slice", func() {
			Avg([]int{})
		})
	})
}
