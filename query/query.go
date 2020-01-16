package query

import (
	"context"
	"time"

	"github.com/object88/shipbot/log"
	"github.com/object88/shipbot/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
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

func (q *Query) Foo() {
	opts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	serverAddr := "localhost:8080"
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		q.Log.Errorf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewShipbotClient(conn)

	req := proto.ListClustersRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = client.ListClusters(ctx, &req)
	if err != nil {
		q.Log.Errorf("Failed to list: %s\n", err.Error())
	}
}
