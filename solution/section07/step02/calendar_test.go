package eventcal_test

import (
	"fmt"
	"sort"
	"testing"
	"time"

	eventcal "github.com/gohandson/toybox-ja/solution/section07/step02"
	"github.com/google/go-cmp/cmp"
)

func TestCalendar_Recent(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		now       func() time.Time
		starts    []time.Time
		durations []time.Duration
		days      int
		want      []int
	}{
		{
			name: "noevents",
			now: func() time.Time {
				return time.Date(2021, 11, 18, 10, 0, 0, 0, time.Local)
			},
			starts:    nil,
			durations: nil,
			days:      1,
			want:      nil,
		},
		{
			name: "recent 1day",
			now: func() time.Time {
				return time.Date(2021, 11, 18, 10, 0, 0, 0, time.Local)
			},
			starts: []time.Time{
				time.Date(2021, 11, 18, 10, 0, 0, 0, time.Local),
				time.Date(2021, 11, 19, 10, 0, 0, 0, time.Local),
			},
			durations: []time.Duration{
				1 * time.Hour,
				1 * time.Hour,
			},
			days: 1,
			want: []int{0},
		},
		{
			name: "recent 2days",
			now: func() time.Time {
				return time.Date(2021, 11, 18, 10, 0, 0, 0, time.Local)
			},
			starts: []time.Time{
				time.Date(2021, 11, 18, 10, 0, 0, 0, time.Local),
				time.Date(2021, 11, 19, 10, 0, 0, 0, time.Local),
			},
			durations: []time.Duration{
				1 * time.Hour,
				1 * time.Hour,
			},
			days: 2,
			want: []int{0, 1},
		},
		{
			name: "past events",
			now: func() time.Time {
				return time.Date(2021, 11, 18, 10, 0, 0, 0, time.Local)
			},
			starts: []time.Time{
				time.Date(2021, 11, 18, 9, 0, 0, 0, time.Local),
				time.Date(2021, 11, 18, 10, 0, 0, 0, time.Local),
			},
			durations: []time.Duration{
				1 * time.Hour,
				1 * time.Hour,
			},
			days: 1,
			want: []int{1},
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cal := eventcal.NewCalendar()
			cal.Now = tt.now

			if len(tt.starts) != len(tt.durations) {
				t.Fatal("startsとdurationsの長さが一致しません")
			}

			for i := range tt.starts {
				cal.Add(&eventcal.Event{
					Title:    fmt.Sprintf("Event %d", i+1),
					Start:    tt.starts[i],
					Duration: tt.durations[i],
				})
			}

			from, got := cal.Recent(tt.days)

			if now := cal.Now(); !from.Equal(now) {
				t.Errorf("from: want %v got %v", now, from)
			}

			gotTitles := make([]string, len(got))
			for i := range got {
				gotTitles[i] = got[i].Title
			}
			sort.Strings(gotTitles)

			wantTitles := make([]string, len(tt.want))
			for i, index := range tt.want {
				wantTitles[i] = fmt.Sprintf("Event %d", index+1)
			}
			sort.Strings(wantTitles)

			if diff := cmp.Diff(wantTitles, gotTitles); diff != "" {
				t.Error(diff)
			}
		})
	}
}
