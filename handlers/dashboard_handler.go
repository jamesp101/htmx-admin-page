package handlers

import (
	"main/templates/pages"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func DashboardHandler() *chi.Mux {
	chi := chi.NewRouter()

	chi.Get("/", rootHandler)

	return chi
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	username := "James"
	templ.Handler(pages.DashboardPage(username)).ServeHTTP(w, r)
}
