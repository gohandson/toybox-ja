package main

import "fmt"

// 変換結果を記録するための構造体
type memory struct {
	converter *converter
	from      float64
	to        float64
}

func (m *memory) String() string {
	return fmt.Sprintf("%.2f[%s] -> %.2f[%s]", m.from, m.converter.fromUnit, m.to, m.converter.toUnit)
}

// 単位変換ツール
type tool struct {
	converters []*converter
	// 履歴を記録するためのフィールド
	memories   []*memory
}

func (t *tool) start() {
	for {

		t.printMemories()

		no, ok := t.inputConverterNo()
		if !ok {
			return
		}

		c := t.converters[no]
		from := t.inputValue(c)
		to := c.convert(from)
		// 履歴として記録する
		m := &memory{
			converter: c,
			from:      from,
			to:        to,
		}
		t.memories = append(t.memories, m)
		// 結果を表示する
		fmt.Println(m.String())

		fmt.Println()
	}
}

func (t *tool) printMemories() {
	if len(t.memories) <= 1 {
		return
	}

	fmt.Println("履歴")
	for _, m := range t.memories {
		fmt.Println(m.String())
	}

	fmt.Println()
}

func (t *tool) inputConverterNo() (int, bool) {
	fmt.Println("以下の単位変換器から選んでください。")
	for i, c := range t.converters {
		fmt.Printf("%d: %s\n", i+1, c.name)
	}
	fmt.Printf("%d: 終了\n", len(t.converters)+1)

	var no int
	for no <= 0 || no > len(t.converters)+1 {
		fmt.Printf("1〜%dの数字>", len(t.converters)+1)
		fmt.Scanln(&no)
	}

	// 終了する
	if no == len(t.converters)+1 {
		return 0, false
	}

	return no - 1, true
}

func (t *tool) inputValue(c *converter) float64 {
	var v float64
	fmt.Println(c.name)
	fmt.Printf("変換する値[%s]>", c.fromUnit)
	fmt.Scanln(&v)
	return v
}
