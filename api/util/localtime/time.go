package localtime

import "time"

// 現在時刻を固定するときに使用
var fixedTime *time.Time = nil

func Now() time.Time {
	if fixedTime != nil {
		return *fixedTime
	}

	return time.Now()
}
