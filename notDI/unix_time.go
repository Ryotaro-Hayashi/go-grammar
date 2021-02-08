package notDI

import "time"

func UnixTimeSample() string {
	unixNow := time.Now().Unix()
	if unixNow %2 != 0 { // Unix時間が奇数
		return "Odd"
	} else {
		return "Even"
	}
}
