package errcode

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	t.Run("create new error successfully", func(t *testing.T) {
		code := 999999
		msg := "这是一个测试错误"
		err := NewError(code, msg)

		assert.NotNil(t, err)
		assert.Equal(t, code, err.Code())
		assert.Equal(t, msg, err.Msg())
		assert.Equal(t, "错误码为：999999 ，错误信息为： 这是一个测试错误", err.Error())
	})

	t.Run("panic on duplicate error code", func(t *testing.T) {
		// Success 已经被定义为 0，再次定义 0 应该 panic
		assert.PanicsWithValue(t, "错误码 0 已经存在，请更换一个", func() {
			NewError(0, "重复的错误码")
		})
	})
}

func TestErrorMethods(t *testing.T) {
	baseErr := NewError(888888, "基础错误: %s")

	t.Run("Code and Msg", func(t *testing.T) {
		assert.Equal(t, 888888, baseErr.Code())
		assert.Equal(t, "基础错误: %s", baseErr.Msg())
	})

	t.Run("Msgf format", func(t *testing.T) {
		newErr := baseErr.Msgf("参数A")
		// 必须是深拷贝，原错误不受影响
		assert.Equal(t, "基础错误: %s", baseErr.Msg())
		assert.Equal(t, "基础错误: 参数A", newErr.Msg())
		assert.Equal(t, 888888, newErr.Code())
	})

	t.Run("Msgr replace", func(t *testing.T) {
		newErr := baseErr.Msgr("完全替换的错误")
		// 必须是深拷贝，原错误不受影响
		assert.Equal(t, "基础错误: %s", baseErr.Msg())
		assert.Equal(t, "完全替换的错误", newErr.Msg())
		assert.Equal(t, 888888, newErr.Code())
	})

	t.Run("WithDetails and Details", func(t *testing.T) {
		assert.Nil(t, baseErr.Details())

		newErr := baseErr.WithDetails("详情1", "详情2")
		// 必须是深拷贝
		assert.Nil(t, baseErr.Details())
		assert.Equal(t, []string{"详情1", "详情2"}, newErr.Details())

		// 测试多次 WithDetails 的行为 (源码实现为覆盖)
		newErr2 := newErr.WithDetails("详情3")
		assert.Equal(t, []string{"详情1", "详情2"}, newErr.Details())
		assert.Equal(t, []string{"详情3"}, newErr2.Details())
	})

	t.Run("WithError and Err", func(t *testing.T) {
		assert.Nil(t, baseErr.Err())

		rawErr := errors.New("底层系统错误")
		newErr := baseErr.WithError(rawErr)

		// 必须是深拷贝
		assert.Nil(t, baseErr.Err())
		assert.Equal(t, rawErr, newErr.Err())
	})

	t.Run("Method chaining", func(t *testing.T) {
		rawErr := errors.New("网络中断")
		finalErr := baseErr.Msgf("测试链式调用").WithDetails("详情A").WithError(rawErr)

		assert.Equal(t, 888888, finalErr.Code())
		assert.Equal(t, "基础错误: 测试链式调用", finalErr.Msg())
		assert.Equal(t, []string{"详情A"}, finalErr.Details())
		assert.Equal(t, rawErr, finalErr.Err())
	})
}

func TestCommonCodeAndModuleCode(t *testing.T) {
	// 简单验证一些包级别的错误码是否被正确初始化
	t.Run("common codes", func(t *testing.T) {
		assert.Equal(t, 0, Success.Code())
		assert.Equal(t, "请求成功", Success.Msg())

		assert.Equal(t, 1, Fail.Code())
		assert.Equal(t, "请求失败", Fail.Msg())

		assert.Equal(t, 100404, NotFound.Code())
	})

	t.Run("module codes", func(t *testing.T) {
		assert.Equal(t, 200000, ErrorUploadFileFail.Code())
		assert.Equal(t, "上传文件失败", ErrorUploadFileFail.Msg())
	})
}
