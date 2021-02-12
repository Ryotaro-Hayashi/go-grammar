package DI

import (
	"testing"
	"time"
)

type MockTimeManager struct {
	MockTime int64
}

func (t *MockTimeManager) NowUnix() int64 {
	return t.MockTime
}

func TestUnixTimeSample(t *testing.T) {
	// テーブル
	cases := []struct {
		t			time.Time
		expected	string
	}{
		{t: time.Unix(1419933531, 0), expected: "Odd"},
		{t: time.Unix(1419933530, 0), expected: "Even"},
	}

	mockTimeManager := &MockTimeManager{}

	for _, c := range cases {
		mockTimeManager.MockTime = c.t.Unix() // Unix時間
		// モックを使うことで（抽象化されて）実際の実装と切り離してテストができる
		result := UnixTimeSample(mockTimeManager) // OddかEven
		if result != c.expected {
			t.Errorf("time is %d and expected is %s but It's %s", mockTimeManager.MockTime, c.expected, result)
		}
	}
}
