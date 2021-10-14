package store

import (
	"context"
	"database/sql"
	"sync"

	"github.com/sanya-spb/Go-Postgres/app/repos/links"
)

var _ links.LinksStore = &Links{}

type Links struct {
	sync.RWMutex
	m map[string]links.TLink
}

func NewLinks() *Links {
	return &Links{
		m: make(map[string]links.TLink),
	}
}

func (link *Links) Create(ctx context.Context, data links.TLink) (string, error) {
	// 	select {
	// 	case <-ctx.Done():
	// 		return "", ctx.Err()
	// 	default:
	// 	}

	// 	id, err := link.GetNextID(ctx)
	// 	if err != nil {
	// 		return "", nil
	// 	}
	// 	data.ID = id

	// 	link.Lock()
	// 	defer link.Unlock()

	// 	link.m[data.ID] = data
	return "", nil
}

func (link *Links) getPerson(ctx context.Context, fName string, lName string) (*links.TLink, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	link.RLock()
	defer link.RUnlock()

	// data, ok := link.m[id]
	// if ok {
	// 	return &data, nil
	// }
	return nil, sql.ErrNoRows
}
