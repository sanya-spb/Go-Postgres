package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sanya-spb/Go-Postgres/app/repos/persons"
)

type Handler struct {
	persons *persons.Persons
}

func NewHandler(persons *persons.Persons) *Handler {
	r := &Handler{
		persons: persons,
	}
	return r
}

type TPerson persons.TPerson

func (hHandler *Handler) GetPerson(ctx context.Context, fName string, lName string) (TPerson, error) {
	if fName == "" {
		return TPerson{}, fmt.Errorf("bad request: fName is empty")
	}

	if lName == "" {
		return TPerson{}, fmt.Errorf("bad request: lName is empty")
	}

	data, err := hHandler.persons.GetPerson(ctx, fName, lName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TPerson{}, ErrLinkNotFound
		}
		return TPerson{}, fmt.Errorf("error when reading: %w", err)
	}

	return TPerson(*data), nil
}
