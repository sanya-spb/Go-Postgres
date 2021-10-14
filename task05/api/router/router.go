package router

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sanya-spb/Go-Postgres/api/handler"
)

type Router struct {
	http.Handler
	hHandler *handler.Handler
}

type TPerson handler.TPerson

func (link *TPerson) Bind(r *http.Request) error {
	return nil
}

func (TPerson) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewRouter(hHandler *handler.Handler) *Router {
	rRouter := &Router{
		hHandler: hHandler,
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	r.Get("/p/{fName}/{lName}", rRouter.GetPerson)

	rRouter.Handler = r
	return rRouter
}

func (rRouter *Router) GetPerson(w http.ResponseWriter, req *http.Request) {
	fName := chi.URLParam(req, "fName")
	lName := chi.URLParam(req, "lName")

	hPerson, err := rRouter.hHandler.GetPerson(req.Context(), fName, lName)
	if err != nil {
		if errors.As(err, &handler.ErrLinkNotFound) {
			render.Render(w, req, Err404(err))
			return
		}
		render.Render(w, req, Err500(err))
		return
	}
	render.Render(w, req, TPerson(hPerson))
}
