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

type TLink handler.TLink

func (link *TLink) Bind(r *http.Request) error {
	return nil
}

func (TLink) Render(w http.ResponseWriter, r *http.Request) error {
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

func (rRouter *Router) Create(w http.ResponseWriter, req *http.Request) {
	link := TLink{}
	if err := render.Bind(req, &link); err != nil {
		render.Render(w, req, Err400(err))
		return
	}
	hLink, err := rRouter.hHandler.Create(req.Context(), handler.TLink(link))
	if err != nil {
		render.Render(w, req, Err500(err))
		return
	}
	render.Status(req, http.StatusCreated)
	render.Render(w, req, TLink(hLink))
}

func (rRouter *Router) GetPerson(w http.ResponseWriter, req *http.Request) {
	fName := chi.URLParam(req, "fName")
	lName := chi.URLParam(req, "lName")

	hLink, err := rRouter.hHandler.GetPerson(req.Context(), fName, lName)
	if err != nil {
		if errors.As(err, &handler.ErrLinkNotFound) {
			render.Render(w, req, Err404(err))
			return
		}
		render.Render(w, req, Err500(err))
		return
	}
	render.Render(w, req, TLink(hLink))
}
