package server

import (
	"context"
	"net/http"
	"time"

	"github.com/sanya-spb/Go-Postgres/app/repos/persons"
)

type Server struct {
	httpServer http.Server
	persons    *persons.Persons
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}

	s.httpServer = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (srv *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	srv.httpServer.Shutdown(ctx)
	cancel()
}

func (srv *Server) Start(persons *persons.Persons) {
	srv.persons = persons
	go srv.httpServer.ListenAndServe()
}

func (srv *Server) Addr() string {
	return srv.httpServer.Addr
}
