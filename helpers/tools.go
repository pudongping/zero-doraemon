package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Choose 是一个简单的三元运算符函数，根据表达式的值返回不同的结果。
func Choose[T any](expr bool, trueVal, falseVal T) T {
	if expr {
		return trueVal
	}
	return falseVal
}

// DeepClone 深拷贝对象
// args:
// src 源对象
// dst 目标对象
//
// return:
// 错误信息
func DeepClone(src, dst any) error {
	srcBytes, err := json.Marshal(src)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("深度拷贝对象时，序列化源对象失败，错误信息为：%+v", err))
	}
	if err := json.Unmarshal(srcBytes, dst); err != nil {
		return errors.Wrap(err, fmt.Sprintf("深度拷贝对象时，反序列化目标对象失败，错误信息为：%+v", err))
	}
	return nil
}
