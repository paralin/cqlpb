package cqlpb

import (
	"github.com/gocql/gocql"
	"github.com/golang/protobuf/proto"
	"github.com/paralin/cqlpb/marshal"
)

type Binding struct {
	err  error
	qry  *gocql.Query
	iter *gocql.Iter
}

func BindQuery(q *gocql.Query) *Binding {
	return &Binding{
		qry: q,
	}
}

func (b *Binding) Close() error {
	if b.err != nil {
		return b.err
	}

	if err := b.iter.Close(); err != nil {
		return err
	}

	return nil
}

func (b *Binding) Scan(dest proto.Message) (bool, error) {
	if b.iter == nil {
		b.iter = b.qry.Iter()
	}

	m := make(map[string]interface{})
	if !b.iter.MapScan(m) {
		return false, nil
	}
	return true, marshal.Unmarshal(dest, m)
}
