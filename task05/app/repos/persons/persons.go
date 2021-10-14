package persons

import (
	"context"
	"fmt"
)

type TPerson struct {
	ID     int    `json:"id"`
	FNname string `json:"fname"`
	LName  string `json:"lname"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}

type PersonsStore interface {
	GetPerson(ctx context.Context, fName string, lName string) (*TPerson, error)
}

type Persons struct {
	store PersonsStore
}

func NewLinks(store PersonsStore) *Persons {
	return &Persons{
		store: store,
	}
}

func (p *Persons) GetPerson(ctx context.Context, fName string, lName string) (*TPerson, error) {
	data, err := p.store.GetPerson(ctx, fName, lName)
	if err != nil {
		return nil, fmt.Errorf("read link error: %w", err)
	}
	return data, nil
}
