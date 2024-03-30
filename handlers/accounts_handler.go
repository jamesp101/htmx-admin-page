package handlers

import (
	"main/templates/pages"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func AccountsHandler() *chi.Mux {

	chi := chi.NewRouter()

	chi.Get("/", accountsPageIndex)
	chi.Get("/{id}", accountInfo)

	return chi
}

func accountsPageIndex(w http.ResponseWriter, r *http.Request) {
	templ.Handler(pages.AccountsPage()).ServeHTTP(w, r)
}

func accountInfo(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "id")
	accountIdStr, err := strconv.Atoi(accountId)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	templ.Handler(pages.AccountInfo(pages.AccountsSample[accountIdStr - 1])).ServeHTTP(w, r)
}
