package eventcal

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type CLI struct {
	Calendar *Calendar
}

func (cli *CLI) Main() int {
	if err := cli.main(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return 1
	}
	return 0
}

func (cli *CLI) main() error {
	for {

		// モードを選択して実行する
		var mode int
		cli.input("[1]イベント入力 [2]直近イベント [3]終了\n>", &mode)

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
			fmt.Println("終了します")
			return nil
		}
	}
}

func (cli *CLI) InputEvents() error {
	var count int
	cli.input("入力するイベント数>", &count)

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

	cli.input("イベント名>", &e.Title)

	var day string
	cli.input("イベント日(YYYYMMDD)>", &day)

	var hmi string
	cli.input("開始時間(HH:mm)>", &hmi)

	var err error
	e.Start, err = time.Parse("2006010215:04", day+hmi)
	if err != nil {
		return nil, err
	}

	var duration string
	cli.input("イベント時間>", &duration)

	e.Duration, err = time.ParseDuration(duration)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (cli *CLI) RecentEvents(days int) error {
	from, events := cli.Calendar.Recent(days)

	if len(events) == 0 {
		fmt.Println("登録されたイベントはありません。")
	}

	if err := cli.printCalendar(from); err != nil {
		return err
	}

	fmt.Println()

	if err := cli.printEvents(events); err != nil {
		return err
	}

	return nil
}

func (cli *CLI) printCalendar(from time.Time) error {
	fmt.Println(from.Format("January 2006"))

	fmt.Println(" Su  Mo  Tu  We  Th  Fr  Sa")

	y, m, _ := from.Date()
	firstday := time.Date(y, m, 1, 0, 0, 0, 0, from.Location())
	lastday := from.AddDate(0, 1, -from.Day())

	fmt.Print(strings.Repeat(" ", 4*int(firstday.Weekday()-time.Sunday)))

	for day := firstday; day.Month() == lastday.Month(); day = day.AddDate(0, 0, 1) {
		switch {
		case day.Day() == from.Day():
			fmt.Printf("*%2d ", day.Day())
		default:
			fmt.Printf(" %2d ", day.Day())
		}

		if day.Weekday() == time.Saturday ||
			day.Day() == lastday.Day() {
			fmt.Println()
		}
	}

	return nil
}

func (cli *CLI) printEvents(events []*Event) error {
	done := make(map[string]bool)
	for _, e := range events {
		day := e.Start.Format("01/02(Mon)")
		if !done[day] {
			fmt.Println("#", day)
			done[day] = true
		}

		fmt.Println("##", e.Title)
		fmt.Println("開始時間:", e.Start.Format("15:04 -"))
		fmt.Println("イベント時間:", e.Duration)
	}

	return nil
}

func (cli *CLI) input(prompt string, v interface{}) {
	fmt.Print(prompt)
	fmt.Scanln(v)
}
