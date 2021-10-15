//go:build integration
// +build integration

package store_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sanya-spb/Go-Postgres/app/repos/persons"
	"github.com/sanya-spb/Go-Postgres/db/postgres/persons/store"
	"github.com/stretchr/testify/require"
)

const (
	DatabaseURL = "postgresql://sanya:passwd@localhost:5432/sauna?sslmode=disable"
)

func TestPersons_GetPerson(t *testing.T) {
	ctx := context.Background()
	p := store.Persons{
		Pool: connect(ctx),
	}
	defer p.Pool.Close()

	tests := []struct {
		fName   string
		lName   string
		prepare func(*pgxpool.Pool)
		check   func(*testing.T, *persons.TPerson, error)
	}{
		{
			fName: "name",
			lName: "surname",
			prepare: func(dbpool *pgxpool.Pool) {
				dbpool.Exec(context.Background(), `insert into personal (id, fname, lname, phone, email) values (100500, 'name', 'surname', '+phone', 'mail@mail.com');`)
			},
			check: func(t *testing.T, result *persons.TPerson, err error) {
				require.NoError(t, err)
				require.Equal(t, result.ID, 100500)
				require.Equal(t, result.FNname, "name")
				require.Equal(t, result.LName, "surname")
				require.Equal(t, result.Phone, "+phone")
				require.Equal(t, result.Email, "mail@mail.com")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.fName, func(t *testing.T) {
			tt.prepare(p.Pool)
			result, err := p.GetPerson(ctx, tt.fName, tt.lName)
			tt.check(t, result, err)
		})
	}
}

func connect(ctx context.Context) *pgxpool.Pool {
	dbpool, err := pgxpool.Connect(ctx, DatabaseURL)
	if err != nil {
		panic(err)
	}

	return dbpool
}
