package main

// 家計簿の項目
type Item struct {
	Category string
	Price    int
}

// 家計簿の処理を行う型
type AccountBook struct {
	items []*Item
}

// 新しいAccountBookを作成する
func NewAccountBook() *AccountBook {
	// TODO: AccountBookのポインタを返す

}

// 新しいItemを追加する
func (ab *AccountBook) AddItem(item *Item) {
	ab.items = append(ab.items, item)
}

// 最近追加したものを最大limit件だけItemを取得する
func (ab *AccountBook) GetItems(limit int) []*Item {
	// TODO: 返す件数を求め変数nに入れる

	if /* TODO: 変数limitより件数が少ない場合 */ {
		n = 0
	}
	return ab.items[n:]
}

// 件数を取得する
func (ab *AccountBook) NumItems() int {
	return len(ab.items)
}
