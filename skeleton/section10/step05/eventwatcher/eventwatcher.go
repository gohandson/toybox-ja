package eventwatcher

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/tenntenn/connpass"
	"github.com/tenntenn/sqlite"
)

type Condition struct {
	ID    int64
	Kind  string
	Value string
}

type EventWatcher struct {
	connpass *connpass.Client
	db       *sql.DB
	mux      *http.ServeMux
	server   *http.Server
}

func New(addr string) (*EventWatcher, error) {
	mux := http.NewServeMux()
	// TODO: ドライバ名にsqlite.DriverName、接続文字列に"eventwatcher.db"を指定してデータベースを開く

	if err != nil {
		return nil, err
	}

	return &EventWatcher{
		connpass: connpass.NewClient(),
		mux:      mux,
		db:       db,
		server:   &http.Server{Addr: addr, Handler: mux},
	}, nil
}

func (ew *EventWatcher) Start() error {
	if err := ew.initDB(context.Background()); err != nil {
		return err
	}
	ew.initHandlers()
	if err := ew.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (ew *EventWatcher) initDB(ctx context.Context) error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS conditions(
		id	INTEGER PRIMARY KEY,
		kind 	TEXT NOT NULL,
		value 	TEXT NOT NULL
	);`

	if _, err := ew.db.ExecContext(ctx, sqlStr); err != nil {
		return err
	}

	return nil
}

func (ew *EventWatcher) Conditions(ctx context.Context, limit int) ([]*Condition, error) {
	const sqlStr = `SELECT id, kind, value FROM conditions LIMIT ?`
	rows, err := ew.db.QueryContext(ctx, sqlStr, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 関数終了時にCloseが呼び出される

	var cs []*Condition
	for rows.Next() {
		var c Condition
		// TODO: ID, Kind, Valueの順でレコードからスキャンする

		if err != nil {
			return nil, err
		}
		cs = append(cs, &c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
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
	const sqlStr = `INSERT INTO conditions(kind, value) VALUES (?,?);`
	r, err := ew.db.ExecContext(ctx, sqlStr, c.Kind, c.Value)
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = id
	return nil
}

func (ew *EventWatcher) RemoveCondition(ctx context.Context, id int64) error {
	const sqlStr = `DELETE FROM conditions WHERE id = ?;`
	_, err := ew.db.ExecContext(ctx, sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}
