package eventcal

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type CLI struct {
	Calendar *Calendar
	Stdout   io.Writer
	Stderr   io.Writer
	Stdin    io.Reader
}

func (cli *CLI) Main() int {
	if err := cli.main(); err != nil {
		fmt.Fprintln(cli.Stderr, "Error:", err)
		return 1
	}
	return 0
}

func (cli *CLI) main() error {
	for {

		// モードを選択して実行する
		var mode int
		if err := cli.input("[1]イベント入力 [2]直近イベント [3]終了\n>", &mode); err != nil {
			return err
		}

		// モードによって処理を変える
		switch mode {
		case 1: // 入力
			if err := cli.InputEvents(); err != nil {
				return err
			}
		case 2: // 直近イベント
			if err := cli.RecentEvents(7); err != nil {
				return err
			}
		case 3: // 終了
			cli.println("終了します")
			return nil
		}
	}
}

func (cli *CLI) InputEvents() error {
	var count int
	if err := cli.input("入力するイベント数>", &count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		e, err := cli.InputEvent()
		if err != nil {
			return err
		}
		cli.Calendar.Add(e)
	}

	return nil
}

func (cli *CLI) InputEvent() (*Event, error) {
	var e Event

	if err := cli.input("イベント名>", &e.Title); err != nil {
		return nil, err
	}

	var day string
	err := cli.input("イベント日(YYYYMMDD)>", &day)
	if err != nil {
		return nil, err
	}
 	
	var hmi string
	err = cli.input("開始時間(HH:mm)>", &hmi)
	if err != nil {
		return nil, err
	}

	e.Start, err = time.Parse("2006010215:04", day+hmi)
	if err != nil {
		return nil, err
	}

	var duration string
	err = cli.input("イベント時間>", &duration)
	if err != nil {
		return nil, err
	}

	e.Duration, err = time.ParseDuration(duration)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (cli *CLI) RecentEvents(days int) error {
	from, events := cli.Calendar.Recent(days)

	if len(events) == 0 {
		if err := cli.println("登録されたイベントはありません。"); err != nil {
			return err
		}
	}

	if err := cli.printCalendar(from); err != nil {
		return err
	}

	if err := cli.println(); err != nil {
		return err
	}

	if err := cli.printEvents(events); err != nil {
		return err
	}

	return nil
}

func (cli *CLI) printCalendar(from time.Time) error {
	if err := cli.println(from.Format("January 2006")); err != nil {
		return err
	}

	if err := cli.println(" Su  Mo  Tu  We  Th  Fr  Sa"); err != nil {
		return err
	}

	y, m, _ := from.Date()
	firstday := time.Date(y, m, 1, 0, 0, 0, 0, from.Location())
	lastday := from.AddDate(0, 1, -from.Day())

	if err := cli.print(strings.Repeat(" ", 4*int(firstday.Weekday()-time.Sunday))); err != nil {
		return err
	}

	for day := firstday; day.Month() == lastday.Month(); day = day.AddDate(0, 0, 1) {
		switch {
		case day.Day() == from.Day():
			if err := cli.printf("*%2d ", day.Day()); err != nil {
				return err
			}
		default:
			if err := cli.printf(" %2d ", day.Day()); err != nil {
				return err
			}
		}

		if day.Weekday() == time.Saturday ||
			day.Day() == lastday.Day() {
			if err := cli.println(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (cli *CLI) printEvents(events []*Event) error {
	done := make(map[string]bool)
	for _, e := range events {
		day := e.Start.Format("01/02(Mon)")
		if !done[day] {
			if err := cli.println("#", day); err != nil {
				return err
			}
			done[day] = true
		}

		if err := cli.println("##", e.Title); err != nil {
			return err
		}

		if err := cli.println("開始時間:", e.Start.Format("15:04 -")); err != nil {
			return err
		}

		if err := cli.println("イベント時間:", e.Duration); err != nil {
			return err
		}
	}

	return nil
}

func (cli *CLI) input(prompt string, v interface{}) error {
	if err := cli.print(prompt); err != nil {
		return err
	}

	if _, err := fmt.Fscanln(cli.Stdin, v); err != nil {
		return err
	}

	return nil
}

func (cli *CLI) print(args ...interface{}) error {
	if _, err := fmt.Fprint(cli.Stdout, args...); err != nil {
		return err
	}
	return nil
}

func (cli *CLI) println(args ...interface{}) error {
	if _, err := fmt.Fprintln(cli.Stdout, args...); err != nil {
		return err
	}
	return nil
}

func (cli *CLI) printf(format string, args ...interface{}) error {
	if _, err := fmt.Fprintf(cli.Stdout, format, args...); err != nil {
		return err
	}
	return nil
}
