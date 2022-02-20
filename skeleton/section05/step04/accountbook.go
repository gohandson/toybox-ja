package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 家計簿の項目
type Item struct {
	Category string
	Price    int
}

// 家計簿の処理を行う型
type AccountBook struct {
	file  string
	items []*Item
}

// 新しいAccountBookを作成する
func NewAccountBook(file string) *AccountBook {
	// AccountBook構造体を作成する
	ab := &AccountBook{
		file: file,
	}

	ab.readItems()

	// AccountBookのポインタを返す
	return ab
}

func (ab *AccountBook) readItems() {
	f, err := os.Open(ab.file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		ss := strings.Split(s.Text(), ",")
		if len(ss) != 2 {
			fmt.Fprintln(os.Stderr, "ファイル形式が不正です")
			os.Exit(1)
		}

		price, err := strconv.Atoi(ss[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "エラー：", err)
			os.Exit(1)
		}

		item := &Item{
			Category: ss[0],
			Price:    price,
		}

		ab.AddItem(item)
	}

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
}

// 新しいItemを追加する
func (ab *AccountBook) AddItem(item *Item) {
	ab.items = append(ab.items, item)

	f, err := os.Create(ab.file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}

	for _, item := range ab.items {
		_, err := fmt.Fprintf(f, "%s,%d\n", item.Category, item.Price)
		if err != nil {
			fmt.Fprintln(os.Stderr, "エラー：", err)
			os.Exit(1)
		}
	}

	if err := f.Close(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
}

// 最近追加したものを最大limit件だけItemを取得する
func (ab *AccountBook) GetItems(limit int) []*Item {
	n := len(ab.items) - limit
	if n < 0 {
		n = 0
	}
	return ab.items[n:]
}

// 件数を取得する
func (ab *AccountBook) NumItems() int {
	return len(ab.items)
}

// 集計結果を取得する
func (ab *AccountBook) GetSummaries() []*Summary {
	m := make(map[string]*Summary)
	var summaries []*Summary

	for _, item := range ab.items {
		s, ok := m[item.Category]
		if !ok {
			s = &Summary{Category: item.Category}
			m[item.Category] = s
			summaries = append(summaries, s)
		}
		s.Count++
		s.Sum += item.Price
	}

	return summaries
}

type Summary struct {
	Category string
	Count    int
	Sum      int
}

// 平均を取得する
func (s *Summary) Avg() float64 {
	// Countが0だとゼロ除算になるため
	// そのまま0を返す
	if s.Count == 0 {
		return 0
	}
	// TODO: 平均を求めて返す

}

func SaveSummary(file string, summaries []*Summary) {

	f, err := os.Create(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}

	cw := csv.NewWriter(f)

	// 品目, 個数, 合計, 平均
	header := []string{"品目", "個数", "合計", "平均"}
	// TODO: 変数headerをCSVとして書き出す

	cw.Flush()

	if /* TODO: 書き込みでエラーが発生していないか確認*/; err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}

	var records [][]string
	for _, s := range summaries {
		records = append(records, []string{
			s.Category,
			strconv.Itoa(s.Count),
			strconv.Itoa(s.Sum),
			// TODO: 平均の少数第2位までを文字列にして最後の要素とする

		})
	}

	if err := /* recordsをCSVとして書き込む */; err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
}
