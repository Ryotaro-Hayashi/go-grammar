package DI

import (
	"fmt"
	"time"
)

// TimeManagerと（テストファイルの）MockTimeManagerがinterfaceを満たす
type ITimeManager interface {
	NowUnix() int64
}

type TimeManager struct {}

// 現在のUnix時間を返す
func (t *TimeManager) NowUnix() int64 {
	return time.Now().Unix()
}

// 引数にinterfaceを渡す
func UnixTimeSample(timeManager ITimeManager) string {
	unixNow := timeManager.NowUnix()
	if unixNow %2 != 0 { // Unix時間が奇数
		fmt.Printf("time is %v and result is Odd\n", unixNow)
		return "Odd"
	} else {
		fmt.Printf("time is %v and result is Even\n", unixNow)
		return "Even"
	}
}

func main() {
	fmt.Println(UnixTimeSample(&TimeManager{}))
}