package eventwatcher

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/tenntenn/connpass"
)

var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

type Condition struct {
	ID    int64  `datastore:"-"`
	Kind  string `datastore:"kind"`
	Value string `datastore:"value"`
}

type EventWatcher struct {
	connpass  *connpass.Client
	datastore *datastore.Client
	mux       *http.ServeMux
	server    *http.Server
}

func New(ctx context.Context, addr string) (*EventWatcher, error) {
	db, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	return &EventWatcher{
		connpass:  connpass.NewClient(),
		datastore: db,
		mux:       mux,
		server:    &http.Server{Addr: addr, Handler: mux},
	}, nil
}

func (ew *EventWatcher) Start() error {
	ew.initHandlers()
	if err := ew.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (ew *EventWatcher) Conditions(ctx context.Context, limit int) ([]*Condition, error) {
	var cs []*Condition
	q := datastore.NewQuery("Condition").Limit(limit)

	keys, err := ew.datastore.GetAll(ctx, q, &cs)
	switch {
	case errors.Is(err, datastore.ErrNoSuchEntity):
		return []*Condition{}, nil
	case err != nil:
		return nil, err
	}

	for i := range cs {
		cs[i].ID = keys[i].ID
	}

	return cs, nil
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

func (ew *EventWatcher) AddCondition(ctx context.Context, c *Condition) error {
	key := datastore.IncompleteKey("Condition", nil)
	newKey, err := ew.datastore.Put(ctx, key, c)
	if err != nil {
		return err
	}
	c.ID = newKey.ID
	return nil
}

func (ew *EventWatcher) RemoveCondition(ctx context.Context, id int64) error {
	key := datastore.IDKey("Condition", id, nil)
	if err := ew.datastore.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}
