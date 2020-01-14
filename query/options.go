package query

import "github.com/object88/shipbot/log"

type Option func(q *Query) error

func SetLogger(l *log.Log) Option {
	return func(q *Query) error {
		q.Log = l
		return nil
	}
}
