package query

import (
	"github.com/object88/shipbot/log"
	"github.com/pkg/errors"
)

type Querier interface{}

type Query struct {
	Log *log.Log
}

func New(opts ...Option) (Querier, error) {
	q := &Query{}

	for _, o := range opts {
		if err := o(q); err != nil {
			return nil, errors.Wrapf(err, "Failed to process options")
		}
	}

	return q, nil
}

func (q *Query) Foo(context string) {

}
