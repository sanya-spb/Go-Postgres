package starter

import (
	"context"
	"log"
	"sync"

	"github.com/sanya-spb/Go-Postgres/app/repos/persons"
)

// application struct
type App struct {
	Persons *persons.Persons
}

// init for App
func NewApp(store persons.PersonsStore) (*App, error) {
	app := &App{
		Persons: persons.NewPersons(store),
	}
	return app, nil
}

type HTTPServer interface {
	Start(p *persons.Persons)
	Stop()
}

// start service
func (app *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs HTTPServer) {
	defer wg.Done()
	hs.Start(app.Persons)
	<-ctx.Done()
	hs.Stop()
}

// print welcome message
func (app *App) Welcome() {
	log.Printf("Let's Go")
}
