package eventwatcher

import (
	"fmt"
	"log"
	"net/http"
)

func (ew *EventWatcher) initHandlers() {
	// TODO: HandleIndexメソッドをパス"/"でmuxフィールドのServeMuxに登録する

}

func (ew *EventWatcher) HandleIndex(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("q")
	if keyword == "" {
		keyword = "golang"
	}
	cs := []*Condition{{Kind: "keyword", Value: keyword}}

	es, err := ew.Events(r.Context(), cs)
	if err != nil {
		ew.error(w, err, http.StatusInternalServerError)
		return
	}

	for _, e := range es {
		if _, err := /* TODO: イベントタイトルをレスポンスとして返す */ ; err != nil {
			ew.error(w, err, http.StatusInternalServerError)
			return
		}
	}
}

func (ew *EventWatcher) error(w http.ResponseWriter, err error, code int) {
	log.Println("Error:", err)
	http.Error(w, http.StatusText(code), code)
}
