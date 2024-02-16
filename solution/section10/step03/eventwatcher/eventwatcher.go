package eventwatcher

import (
	"context"
	"net/http"
	"time"

	"github.com/tenntenn/connpass"
)

type Condition struct {
	Kind  string
	Value string
}

type EventWatcher struct {
	connpass *connpass.Client
	mux      *http.ServeMux
	server   *http.Server
}

func New(addr string) (*EventWatcher, error) {
	mux := http.NewServeMux()

	return &EventWatcher{
		connpass: connpass.NewClient(),
		mux:      mux,
		server:   &http.Server{Addr: addr, Handler: mux},
	}, nil
}

func (ew *EventWatcher) Start() error {
	ew.initHandlers()
	if err := ew.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (ew *EventWatcher) Events(ctx context.Context, cs []*Condition) ([]*connpass.Event, error) {
	params, err := ew.makeParams(cs)
	if err != nil {
		return nil, err
	}

	urlValues, err := connpass.SearchParam(params...)
	if err != nil {
		return nil, err
	}

	r, err := ew.connpass.Search(ctx, urlValues)
	if err != nil {
		return nil, err
	}

	return r.Events, nil
}

func (ew *EventWatcher) makeParams(cs []*Condition) ([]connpass.Param, error) {
	params := make([]connpass.Param, len(cs))
	for i := range cs {
		switch cs[i].Kind {
		case "keyword":
			params[i] = connpass.Keyword(cs[i].Value)
		case "keyword_or":
			params[i] = connpass.KeywordOr(cs[i].Value)
		case "ym":
			tm, err := time.Parse("200601", cs[i].Value)
			if err != nil {
				return nil, err
			}
			params[i] = connpass.YearMonth(tm.Year(), tm.Month())
		case "ymd":
			tm, err := time.Parse("20060102", cs[i].Value)
			if err != nil {
				return nil, err
			}
			params[i] = connpass.YearMonthDay(tm.Year(), tm.Month(), tm.Day())
		}
	}

	return params, nil
}
