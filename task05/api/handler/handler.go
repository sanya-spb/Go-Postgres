package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sanya-spb/Go-Postgres/app/repos/links"
)

type Handler struct {
	links *links.Links
}

func NewHandler(links *links.Links) *Handler {
	r := &Handler{
		links: links,
	}
	return r
}

type TLink links.TLink

func (hHandler *Handler) GetPerson(ctx context.Context, fName string, lName string) (TLink, error) {
	if fName == "" {
		return TLink{}, fmt.Errorf("bad request: fName is empty")
	}

	if lName == "" {
		return TLink{}, fmt.Errorf("bad request: lName is empty")
	}

	data, err := hHandler.links.GetPerson(ctx, fName, lName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TLink{}, ErrLinkNotFound
		}
		return TLink{}, fmt.Errorf("error when reading: %w", err)
	}

	return TLink(*data), nil
}
