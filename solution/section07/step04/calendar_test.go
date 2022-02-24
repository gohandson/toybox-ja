package eventcal_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	eventcal "github.com/gohandson/toybox-ja/solution/section07/step04"
)

func TestCalendar_Recent(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		now       string
		starts    string
		durations string
		days      int
		want      []int
	}{
		{"noevents", "2021/11/18 10:00", "", "", 1, nil},
		{"recent 1day", "2021/11/18 10:00", "2021/11/18 10:00, 2021/11/19 10:00", "1h, 1h", 1, []int{0}},
		{"recent 2days", "2021/11/18 10:00", "2021/11/18 10:00, 2021/11/19 10:00", "1h, 1h", 2, []int{0, 1}},
		{"past events", "2021/11/18 10:00", "2021/11/18 09:00, 2021/11/18 10:00, 2021/11/19 10:00", "1h, 1h, 1h", 1, []int{1}},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cal := eventcal.NewCalendar()
			cal.Clock = fakeClock(t, tt.now)

			var starts []time.Time
			if tt.starts != "" {
				starts = dates(t, strings.Split(tt.starts, ",")...)
			}

			var ds []time.Duration
			if tt.durations != "" {
				ds = durations(t, strings.Split(tt.durations, ",")...)
			}

			if len(starts) != len(ds) {
				t.Fatal("startsとdurationsの長さが一致しません")
			}

			for i := range starts {
				cal.Add(&eventcal.Event{
					Title:    fmt.Sprintf("Event %d", i+1),
					Start:    starts[i],
					Duration: ds[i],
				})
			}

			from, got := cal.Recent(tt.days)

			if now := cal.Clock.Now(); !from.Equal(now) {
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

func fakeClock(t *testing.T, tmstr string) eventcal.Clock {
	t.Helper()
	tm := date(t, tmstr)
	return eventcal.ClockFunc(func() time.Time {
		return tm
	})
}

func dates(t *testing.T, tmstr ...string) []time.Time {
	t.Helper()
	tms := make([]time.Time, len(tmstr))
	for i := range tmstr {
		tms[i] = date(t, tmstr[i])
	}
	return tms
}

func date(t *testing.T, tmstr string) time.Time {
	t.Helper()

	tm, err := time.Parse("2006/01/02 15:04", strings.TrimSpace(tmstr))
	if err != nil {
		t.Fatal("予想外のエラー:", err)
	}

	return tm
}

func durations(t *testing.T, dstr ...string) []time.Duration {
	t.Helper()

	ds := make([]time.Duration, len(dstr))
	for i := range dstr {
		ds[i] = duration(t, dstr[i])
	}
	return ds
}

func duration(t *testing.T, dstr string) time.Duration {
	t.Helper()

	d, err := time.ParseDuration(strings.TrimSpace(dstr))
	if err != nil {
		t.Fatal("予想外のエラー:", err)
	}

	return d
}
