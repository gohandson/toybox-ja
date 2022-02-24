package eventcal

import (
	"sort"
	"time"
)

// Eventは1つのイベント（勉強会）を表す
type Event struct {
	Title    string        // イベントのタイトル
	Start    time.Time     // 開始時間
	// TODO: イベントの時間（Durationフィールド）を時間間隔を表す型で定義

}

// Calendarはイベントカレンダーを表す
type Calendar struct {
	events []*Event
}

func NewCalendar() *Calendar {
	return new(Calendar)
}

func (cal *Calendar) Add(e *Event) {
	cal.events = append(cal.events, e)
}

// 近日開催されるイベントを取得する
func (cal *Calendar) Recent(days int) (time.Time, []*Event) {
	var recents []*Event
	// TODO: 現在時刻を取得して変数fromに代入

	// TODO: 現在時刻にdays日分足して、時刻を0時0分にする


	for _, e := range cal.events {
		if e.Start.Equal(from) || (/* TODO: 現在時刻より後か */ && e.Start.Before(to)) {
			recents = append(recents, e)
		}
	}

	sort.Slice(recents, func(i, j int) bool {
		return recents[i].Start.Before(recents[j].Start)
	})

	return from, recents
}
