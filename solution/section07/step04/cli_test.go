package eventcal_test

import (
	"bytes"
	"flag"
	"strings"
	"testing"

	eventcal "github.com/gohandson/toybox-ja/solution/section07/step04"
	"github.com/tenntenn/golden"
)

var (
	flagUpdate bool
)

func init() {
	flag.BoolVar(&flagUpdate, "update", false, "update golden files")
}

func TestCLI_Main(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name         string
		now          string
		input        string
		wantError    bool
		wantExitCode int
	}{
		{"no event", "2021/11/18 10:00", "2 3", false, 0},
		{"input one event", "2021/11/18 10:00", "1 1 Event1 20211118 10:00 1h 2 3", false, 0},
		{"input two event", "2021/11/18 10:00", "1 2 Event1 20211118 10:00 1h Event2 20211118 12:00 1h 2 3", false, 0},
		{"past event", "2021/11/18 10:00", "1 2 Event1 20211118 09:00 1h Event2 20211118 10:00 1h 2 3", false, 0},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cal := eventcal.NewCalendar()
			cal.Clock = fakeClock(t, tt.now)

			var stdout, stderr bytes.Buffer
			input := strings.NewReader(strings.ReplaceAll(tt.input, " ", "\n"))

			cli := &eventcal.CLI{
				Calendar: cal,
				Stdout:   &stdout,
				Stderr:   &stderr,
				Stdin:    input,
			}

			code := cli.Main()

			errmsg := stderr.String()
			switch {
			case !tt.wantError && errmsg != "":
				t.Fatal("予期せぬエラー:", errmsg)
			case tt.wantError && errmsg == "":
				t.Fatal("期待するエラーが発生しませんでした")
			}

			if code != tt.wantExitCode {
				t.Fatal("予期せぬ終了コード:", code)
			}

			name := strings.ReplaceAll(tt.name, " ", "-")

			if flagUpdate {
				golden.Update(t, "testdata", name, &stdout)
				t.Skip()
			}

			if diff := golden.Diff(t, "testdata", name, &stdout); diff != "" {
				t.Error(diff)
			}
		})
	}
}
