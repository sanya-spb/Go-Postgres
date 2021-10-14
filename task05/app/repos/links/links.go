package links

import (
	"context"
	"fmt"
)

type TLink struct {
	id    int
	fname string
	lname string
	phone string
	email string
	// ID string `json:"id"`
	// Name      string    `json:"name"`
	// URL       string    `json:"url"`
	// Descr     string    `json:"descr"`
	// CreatedAt time.Time `json:"created_at"`
	// DeleteAt  time.Time `json:"delete_at"`
	// User      string    `json:"user"`
	// GoCount   int       `json:"go_count"`
}

type LinksStore interface {
	GetPerson(ctx context.Context, fName string, lName string) (*TLink, error)
}

type Links struct {
	store LinksStore
}

func NewLinks(store LinksStore) *Links {
	return &Links{
		store: store,
	}
}

func (link *Links) GetPerson(ctx context.Context, fName string, lName string) (*TLink, error) {
	data, err := link.store.GetPerson(ctx, fName, lName)
	if err != nil {
		return nil, fmt.Errorf("read link error: %w", err)
	}
	return data, nil
}
