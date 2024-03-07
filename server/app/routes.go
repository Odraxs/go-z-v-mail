package app

import (
	"net/http"

	"github.com/Odraxs/go-z-v-mail/server/email_search"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/emailSearch", app.loadEmailSearchRoutes)

	app.router = router
}

func (app *App) loadEmailSearchRoutes(router chi.Router) {
	emailSearchHandler := &email_search.EmailSearchHandler{
		Repo: email_search.NewZincsearchRepository(),
	}

	router.Post("/", emailSearchHandler.Search)
}
