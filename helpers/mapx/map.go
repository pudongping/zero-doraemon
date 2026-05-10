package mapx

// Keys 获取 map 的所有键，并作为一个切片返回
// 注意：返回的键的顺序是不确定的
//
// 示例:
// Keys(map[string]int{"a": 1, "b": 2})
// 返回: []string{"a", "b"} (顺序随机)
func Keys[K comparable, V any](m map[K]V) []K {
	if len(m) == 0 {
		return make([]K, 0)
	}

	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values 获取 map 的所有值，并作为一个切片返回
// 注意：返回的值的顺序是不确定的
//
// 示例:
// Values(map[string]int{"a": 1, "b": 2})
// 返回: []int{1, 2} (顺序随机)
func Values[K comparable, V any](m map[K]V) []V {
	if len(m) == 0 {
		return make([]V, 0)
	}

	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Merge 合并多个 map，返回一个新 map（不改变原 map）
// 遇到相同的键时，后面 map 的值会覆盖前面 map 的值
//
// 示例:
// Merge(map[string]int{"a": 1}, map[string]int{"a": 2, "b": 3})
// 返回: map[string]int{"a": 2, "b": 3}
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	// 预估一个初始容量，减少扩容开销
	var totalLen int
	for _, m := range maps {
		totalLen += len(m)
	}

	if totalLen == 0 {
		return make(map[K]V)
	}

	result := make(map[K]V, totalLen)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// Filter 过滤 map，返回满足 predicate 条件的键值对组成的新 map
//
// 示例:
// Filter(map[string]int{"a": 1, "b": 2, "c": 3}, func(k string, v int) bool { return v > 1 })
// 返回: map[string]int{"b": 2, "c": 3}
func Filter[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	// 不预先分配 len(m) 容量的原因是为了防止“内存浪费/泄漏”，同 arrayx.Filter
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// MapValues 对 map 中的每个值执行 iteratee 操作，并返回一个值类型为新类型的新 map，键保持不变
//
// 示例:
// 假设 strconv.Itoa 可以直接作为闭包传入
// MapValues(map[string]int{"a": 1, "b": 2}, func(k string, v int) string { return strconv.Itoa(v) })
// 返回: map[string]string{"a": "1", "b": "2"}
func MapValues[K comparable, V any, R any](m map[K]V, iteratee func(K, V) R) map[K]R {
	if len(m) == 0 {
		return make(map[K]R)
	}

	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = iteratee(k, v)
	}
	return result
}

// MapKeys 对 map 中的每个键执行 iteratee 操作，并返回一个键类型为新类型的新 map，值保持不变
// 注意：如果 iteratee 生成了重复的键，后面的值会覆盖前面的值
//
// 示例:
// MapKeys(map[string]int{"a": 1, "b": 2}, func(k string, v int) string { return k + "_new" })
// 返回: map[string]int{"a_new": 1, "b_new": 2}
func MapKeys[K comparable, V any, R comparable](m map[K]V, iteratee func(K, V) R) map[R]V {
	if len(m) == 0 {
		return make(map[R]V)
	}

	result := make(map[R]V, len(m))
	for k, v := range m {
		result[iteratee(k, v)] = v
	}
	return result
}

// Has 检查 map 中是否存在指定的键
//
// 示例:
// Has(map[string]int{"a": 1}, "a")
// 返回: true
func Has[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// Pick 提取 map 中指定的键值对，返回一个新 map
// 如果指定的键在原 map 中不存在，则会被忽略
//
// 示例:
// Pick(map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "c"})
// 返回: map[string]int{"a": 1, "c": 3}
func Pick[K comparable, V any](m map[K]V, keys []K) map[K]V {
	result := make(map[K]V, len(keys))
	for _, key := range keys {
		if val, ok := m[key]; ok {
			result[key] = val
		}
	}
	return result
}

// Omit 排除 map 中指定的键，返回包含剩余键值对的新 map
//
// 示例:
// Omit(map[string]int{"a": 1, "b": 2, "c": 3}, []string{"b"})
// 返回: map[string]int{"a": 1, "c": 3}
func Omit[K comparable, V any](m map[K]V, keys []K) map[K]V {
	// 将需要排除的 keys 转换为 map，提高查找效率 (O(1))
	omitKeys := make(map[K]struct{}, len(keys))
	for _, key := range keys {
		omitKeys[key] = struct{}{}
	}

	result := make(map[K]V)
	for k, v := range m {
		if _, shouldOmit := omitKeys[k]; !shouldOmit {
			result[k] = v
		}
	}
	return result
}

// Clone 浅拷贝一个 map (即只拷贝一层的键值对，如果 value 是指针或引用类型，仍然会指向同一内存)
//
// 示例:
// Clone(map[string]int{"a": 1}) // 返回: map[string]int{"a": 1} 的一个新实例
func Clone[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}

	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// Invert 键值反转：将 map 的键和值互换，返回一个新的 map
// 注意事项：
// 1. 由于反转后原 map 的值（Value）变成了新 map 的键（Key），所以原 map 的值类型 V 也必须是 comparable（可比较的）。
// 2. 如果原 map 中存在多个相同的 Value，由于 Go 语言 map 遍历的无序性，反转后被保留的 Key 是随机的（后面的会覆盖前面的）。
//
// 示例:
// Invert(map[string]int{"a": 1, "b": 2})
// 返回: map[int]string{1: "a", 2: "b"}
func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	if len(m) == 0 {
		return make(map[V]K)
	}

	result := make(map[V]K, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}
