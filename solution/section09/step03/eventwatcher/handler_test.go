package eventwatcher_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gohandson/toybox-ja/solution/section09/step03/eventwatcher"
)

func TestEventWatcher_HandleIndex(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/?q=golang", nil)
	ew, err := eventwatcher.New(":8080")
	if err != nil {
		t.Fatal("予期せぬエラー:", err)
	}
	ew.HandleIndex(w, r)

	res := w.Result()
	t.Cleanup(func() {
		res.Body.Close()
	})
	if res.StatusCode != http.StatusOK {
		t.Error("期待しないステータスコード:", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal("予期せぬエラー:", err)
	}

	// 改行区切り
	n := len(bytes.Split(body, []byte("\n")))
	if len(body) == 0 || n <= 0 {
		t.Error("1件以上の結果が表示されることを期待していた")
	}

	t.Log(string(body))
}
