package links

import (
	"context"
	"fmt"
)

type TLink struct {
	ID string `json:"id"`
	// Name      string    `json:"name"`
	// URL       string    `json:"url"`
	// Descr     string    `json:"descr"`
	// CreatedAt time.Time `json:"created_at"`
	// DeleteAt  time.Time `json:"delete_at"`
	// User      string    `json:"user"`
	// GoCount   int       `json:"go_count"`
}

type LinksStore interface {
	Create(ctx context.Context, data TLink) (string, error)
	getPerson(ctx context.Context, fName string, lName string) (*TLink, error)
}

type Links struct {
	store LinksStore
}

func NewLinks(store LinksStore) *Links {
	return &Links{
		store: store,
	}
}

// Create new Link with returning it
func (link *Links) Create(ctx context.Context, data TLink) (*TLink, error) {
	id, err := link.store.Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("create link error: %w", err)
	}
	data.ID = id
	return &data, nil
}

// Return Link by ID
func (link *Links) getPerson(ctx context.Context, fName string, lName string) (*TLink, error) {
	data, err := link.store.getPerson(ctx, fName, lName)
	if err != nil {
		return nil, fmt.Errorf("read link error: %w", err)
	}
	return data, nil
}
