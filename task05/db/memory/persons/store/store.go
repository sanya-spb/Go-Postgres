package store

import (
	"context"
	"database/sql"
	"sync"

	"github.com/sanya-spb/Go-Postgres/app/repos/persons"
)

var _ persons.PersonsStore = &Persons{}

type Persons struct {
	sync.RWMutex
	m map[string]persons.TPerson
}

func NewPersons() *Persons {
	return &Persons{
		m: make(map[string]persons.TPerson),
	}
}

func (p *Persons) GetPerson(ctx context.Context, fName string, lName string) (*persons.TPerson, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	p.RLock()
	defer p.RUnlock()

	// data, ok := link.m[id]
	// if ok {
	// 	return &data, nil
	// }
	return nil, sql.ErrNoRows
}
