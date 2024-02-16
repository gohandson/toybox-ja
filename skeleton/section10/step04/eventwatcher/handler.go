package eventwatcher

import (
	"html/template"
	"log"
	"net/http"

	"github.com/tenntenn/connpass"
)

var (
	// TODO: eventwatcher/_template/*.htmlをテンプレートとしてパースする

)

func (ew *EventWatcher) initHandlers() {
	ew.mux.HandleFunc("/", ew.HandleIndex)
}

func (ew *EventWatcher) HandleIndex(w http.ResponseWriter, r *http.Request) {
	cs := []*Condition{{Kind: "keyword", Value: "golang"}}
	es, err := ew.Events(r.Context(), cs)
	if err != nil {
		ew.error(w, err, http.StatusInternalServerError)
		return
	}

	data := struct {
		Conditions []*Condition
		Events     []*connpass.Event
	}{
		Conditions: cs,
		Events:     es,
	}

	if err := tmpl.ExecuteTemplate(/* TODO: 結果をレスポンスとして返す */, "index", data); err != nil {
		ew.error(w, err, http.StatusInternalServerError)
		return
	}
}

func (ew *EventWatcher) error(w http.ResponseWriter, err error, code int) {
	log.Println("Error:", err)
	http.Error(w, http.StatusText(code), code)
}
