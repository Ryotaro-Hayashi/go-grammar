package notDI

import (
	"fmt"
	"testing"
)

// これだとテスト結果が実行時の時間に左右されてしまう
func TestUnixTimeSample(t *testing.T) {
	expected := "Odd"
	result := UnixTimeSample()
	if result != expected {
		t.Errorf("It's not %s", expected)
	}
	fmt.Println("result is", result)
}
