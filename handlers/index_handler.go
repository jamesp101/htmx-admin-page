package handlers

import (
	"main/services"
	"main/templates/pages"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func IndexHandler() *chi.Mux {
	chi := chi.NewRouter()

	chi.Get("/", indexPageHandler)

	return chi
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {

	services.GetSession(r, "session")
	templ.Handler(pages.IndexPage()).ServeHTTP(w, r)
}
