package arrayx

// In 判断元素是否存在于切片中
//
// 示例:
// In(2, []int{1, 2, 3}) // 返回: true
// In(4, []int{1, 2, 3}) // 返回: false
func In[T comparable](element T, slice []T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

// Unique 去重，保持切片中元素原有的顺序
//
// 示例:
// Unique([]int{1, 2, 2, 3, 1})
// 返回: []int{1, 2, 3}
func Unique[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return make([]T, 0)
	}

	seen := make(map[T]struct{}, len(slice))
	// 不预先分配 len(slice) 容量，防止极端重复情况下（例如 100万个相同的元素）产生的内存浪费
	result := make([]T, 0)
	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Merge 合并多个切片，返回一个新切片（不改变原切片）
//
// 示例:
// Merge([]int{1, 2}, []int{3, 4}, []int{5})
// 返回: []int{1, 2, 3, 4, 5}
func Merge[T any](slices ...[]T) []T {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}

	if totalLen == 0 {
		return make([]T, 0)
	}

	result := make([]T, 0, totalLen)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// Diff 差集：返回存在于 slice1 但不存在于 slice2 的元素（保持 slice1 中原有顺序）
//
// 示例:
// Diff([]int{1, 2, 3, 4}, []int{2, 4, 5})
// 返回: []int{1, 3}
func Diff[T comparable](slice1, slice2 []T) []T {
	if len(slice1) == 0 {
		return make([]T, 0)
	}
	if len(slice2) == 0 {
		result := make([]T, len(slice1))
		copy(result, slice1)
		return result
	}

	seen := make(map[T]struct{}, len(slice2))
	for _, v := range slice2 {
		seen[v] = struct{}{}
	}

	// 不预先分配 len(slice1) 容量，防止大部分元素被差集过滤掉时产生的内存浪费
	result := make([]T, 0)
	for _, v := range slice1 {
		if _, ok := seen[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

// Intersect 交集：返回同时存在于 slice1 和 slice2 的元素（去重且保持 slice1 的相对顺序）
//
// 示例:
// Intersect([]int{1, 2, 3, 2}, []int{2, 4, 3})
// 返回: []int{2, 3}
func Intersect[T comparable](slice1, slice2 []T) []T {
	if len(slice1) == 0 || len(slice2) == 0 {
		return make([]T, 0)
	}

	seen := make(map[T]struct{}, len(slice2))
	for _, v := range slice2 {
		seen[v] = struct{}{}
	}

	// 交集结果不可能超过 slice1 的长度，但通常远小于 slice1，
	// 此处不分配满容量是为了避免空间浪费。
	result := make([]T, 0)
	added := make(map[T]struct{})
	for _, v := range slice1 {
		if _, ok := seen[v]; ok {
			if _, alreadyAdded := added[v]; !alreadyAdded {
				added[v] = struct{}{}
				result = append(result, v)
			}
		}
	}
	return result
}

// Filter 过滤切片，返回满足 predicate 条件的元素组成的新切片
//
// 示例:
// Filter([]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 })
// 返回: []int{2, 4}
func Filter[T any](slice []T, predicate func(T) bool) []T {
	// 不预先分配 len(slice) 容量的原因是为了防止“内存浪费/泄漏”。
	// 假如原切片有 100 万数据，而过滤后只有 1 个满足条件，如果分配满容量，
	// 返回的切片将一直占用 100 万个元素的内存。使用 make([]T, 0) 虽然会有扩容消耗，但空间更安全。
	result := make([]T, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map 对切片中的每个元素执行 iteratee 操作，并返回一个新类型的新切片
//
// 示例:
// 假设 strconv.Itoa 可以直接作为闭包传入
// Map([]int{1, 2, 3}, func(v int) string { return fmt.Sprintf("%d", v) })
// 返回: []string{"1", "2", "3"}
func Map[T any, R any](slice []T, iteratee func(T) R) []R {
	if len(slice) == 0 {
		return make([]R, 0)
	}

	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = iteratee(v)
	}
	return result
}

// Reverse 反转切片中的元素顺序
// 此方法会直接修改原切片 (in-place)，并返回反转后的原切片以便于链式调用
//
// 示例:
// Reverse([]int{1, 2, 3, 4}) // 返回: []int{4, 3, 2, 1}
func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// Chunk 将切片按照指定大小切割成多个小的切片组
//
// 示例:
// Chunk([]int{1, 2, 3, 4, 5}, 2)
// 返回: [][]int{{1, 2}, {3, 4}, {5}}
func Chunk[T any](slice []T, size int) [][]T {
	if len(slice) == 0 || size <= 0 {
		return make([][]T, 0)
	}

	chunks := make([][]T, 0, (len(slice)+size-1)/size)
	for size < len(slice) {
		slice, chunks = slice[size:], append(chunks, slice[0:size:size])
	}
	chunks = append(chunks, slice)

	return chunks
}

// ContainsBy 检查切片中是否包含满足 predicate 条件的元素
//
// 示例:
// ContainsBy([]int{1, 2, 3}, func(v int) bool { return v > 2 })
// 返回: true
func ContainsBy[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// GroupBy 根据 iteratee 函数的返回值对切片元素进行分组
// 非常适用于将数据库查出的列表按某个字段（如状态、分类）进行分组
//
// 示例:
//
//	GroupBy([]int{1, 2, 3, 4}, func(v int) string {
//	    if v%2 == 0 { return "even" }
//	    return "odd"
//	})
//
// 返回: map[string][]int{"odd": {1, 3}, "even": {2, 4}}
func GroupBy[T any, K comparable](slice []T, iteratee func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := iteratee(v)
		result[key] = append(result[key], v)
	}
	return result
}

// KeyBy 根据 iteratee 函数的返回值，将切片转换为 Map。
// 注意：如果有多个元素返回相同的键，后面的元素会覆盖前面的元素。
// 非常适用于将数据库查出的列表转换为以 ID 为 Key 的 Map，方便快速通过 ID 查找数据。
//
// 示例:
// type User struct { ID int; Name string }
// users := []User{{1, "Alice"}, {2, "Bob"}}
// KeyBy(users, func(u User) int { return u.ID })
// 返回: map[int]User{1: {1, "Alice"}, 2: {2, "Bob"}}
func KeyBy[T any, K comparable](slice []T, iteratee func(T) K) map[K]T {
	result := make(map[K]T, len(slice))
	for _, v := range slice {
		result[iteratee(v)] = v
	}
	return result
}

// Clone 浅拷贝切片（仅拷贝切片的第一层元素）
// 如果切片中的元素是指针或引用类型，它们依然会指向同一块内存
//
// 示例:
// Clone([]int{1, 2, 3})
// 返回: []int{1, 2, 3} 的一个新实例
func Clone[T any](slice []T) []T {
	if slice == nil {
		return nil
	}
	result := make([]T, len(slice))
	copy(result, slice)
	return result
}

// Number 定义了支持算术运算的泛型约束
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Max 返回切片中的最大值
// 如果切片为空，将抛出 panic
//
// 示例:
// Max([]int{1, 5, 3})
// 返回: 5
func Max[T Number](slice []T) T {
	if len(slice) == 0 {
		panic("arrayx: Max called on an empty slice")
	}

	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

// Min 返回切片中的最小值
// 如果切片为空，将抛出 panic
//
// 示例:
// Min([]int{1, 5, 3})
// 返回: 1
func Min[T Number](slice []T) T {
	if len(slice) == 0 {
		panic("arrayx: Min called on an empty slice")
	}

	min := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

// Sum 返回切片中所有元素的总和
// 如果切片为空，将返回 0 (该类型的零值)
//
// 示例:
// Sum([]int{1, 2, 3})
// 返回: 6
func Sum[T Number](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Avg 返回切片中所有元素的平均值，返回类型为 float64
// 如果切片为空，将抛出 panic
//
// 示例:
// Avg([]int{1, 2, 3, 4})
// 返回: 2.5
func Avg[T Number](slice []T) float64 {
	if len(slice) == 0 {
		panic("arrayx: Avg called on an empty slice")
	}

	var sum float64
	for _, v := range slice {
		sum += float64(v)
	}
	return sum / float64(len(slice))
}
