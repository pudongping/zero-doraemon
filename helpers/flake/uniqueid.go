package flake

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var sonyFlake *sonyflake.Sonyflake

func init() {
	var (
		st  time.Time
		err error
	)

	startTime := "2026-01-01" // 初始化一个开始的时间，表示从这个时间开始算起
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		panic(fmt.Sprintf("init sony_flake error: %+v", err))
	}

	settings := sonyflake.Settings{
		StartTime: st,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	if sonyFlake == nil {
		panic("sony_flake not created")
	}
}

// GenUniqueID 雪花算法生成唯一分布式ID
func GenUniqueID() int64 {
	id, err := sonyFlake.NextID()
	if err != nil {
		panic(fmt.Sprintf("gen unique id error: %+v", err))
	}

	return int64(id)
}
