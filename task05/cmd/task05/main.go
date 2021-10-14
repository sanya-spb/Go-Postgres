package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/sanya-spb/Go-Postgres/api/handler"
	"github.com/sanya-spb/Go-Postgres/api/router"
	"github.com/sanya-spb/Go-Postgres/api/server"
	"github.com/sanya-spb/Go-Postgres/app/repos/persons"
	"github.com/sanya-spb/Go-Postgres/app/starter"
	"github.com/sanya-spb/Go-Postgres/db/memory/persons/store"
)

func main() {
	store := store.NewLinks()
	app, err := starter.NewApp(store)
	if err != nil {
		log.Fatalln(err.Error())
	}

	app.Welcome()

	persons := persons.NewLinks(store)
	appHandler := handler.NewHandler(persons)
	appRouter := router.NewRouter(appHandler)
	appServer := server.NewServer(":8080", appRouter)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	log.Printf("listen at: %s\n", appServer.Addr())
	go app.Serve(ctx, wg, appServer)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
