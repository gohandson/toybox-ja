package eventcal

import (
	"sort"
	"time"
)

// Eventは1つのイベント（勉強会）を表す
type Event struct {
	Title    string        // イベントのタイトル
	Start    time.Time     // 開始時間
	Duration time.Duration // イベントの時間
}

// Calendarはイベントカレンダーを表す
type Calendar struct {
	Clock  Clock
	events []*Event
}

func NewCalendar() *Calendar {
	return &Calendar{
		Clock: DefaultClock,
	}
}

func (cal *Calendar) Add(e *Event) {
	cal.events = append(cal.events, e)
}

// 近日開催されるイベントを取得する
func (cal *Calendar) Recent(days int) (time.Time, []*Event) {
	var recents []*Event
	// 現在時刻を取得
	from := cal.Clock.Now()
	// 取得する範囲の終了日時
	to := from.AddDate(0, 0, days).Truncate(24 * time.Hour)

	for _, e := range cal.events {
		if e.Start.Equal(from) || (e.Start.After(from) && e.Start.Before(to)) {
			recents = append(recents, e)
		}
	}

	sort.Slice(recents, func(i, j int) bool {
		return recents[i].Start.Before(recents[j].Start)
	})

	return from, recents
}
