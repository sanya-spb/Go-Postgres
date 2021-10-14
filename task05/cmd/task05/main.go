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
	"github.com/sanya-spb/Go-Postgres/app/repos/links"
	"github.com/sanya-spb/Go-Postgres/app/starter"
	"github.com/sanya-spb/Go-Postgres/db/memory/links/store"
)

func main() {
	store := store.NewLinks()
	app, err := starter.NewApp(store)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// if _, err := os.Stat(filepath.Dir(app.Config.LogAccess)); os.IsNotExist(err) {
	// 	log.Fatalln(err.Error())
	// }
	// if _, err := os.Stat(filepath.Dir(app.Config.LogErrors)); os.IsNotExist(err) {
	// 	log.Fatalln(err.Error())
	// }

	app.Welcome()

	links := links.NewLinks(store)
	appHandler := handler.NewHandler(links)
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
