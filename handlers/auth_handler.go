package handlers

import (
	"log"
	"main/services"
	"main/templates/components"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func AuthHandler() *chi.Mux {
	chi := chi.NewRouter()

	chi.Post("/login", handleLogin)
	chi.Post("/logout", handleLogout)

	return chi
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	log.Printf("Attempting to login from %s", r.RemoteAddr)

	username := r.FormValue("username")
	password := r.FormValue("password")

	if !(username == "admin" && password == "admin") {
		templ.Handler(components.LoginForm(components.LoginFormData{
			Username: username,
			Password: password,
			Error:    "Invalid username or password",
		})).ServeHTTP(w, r)

		log.Printf("Login fail for %s \n", username)
		return
	}

	_, err := services.CreateSession(services.CreateSessionRequest{
		Username: username,
		R:        r,
		W:        w,
	})
	if err != nil {
		log.Fatalf("Unable to create session: %s", err.Error())
	}

	// services.CreateSessionCookie(w, sessionId.Name())


	w.Header().Add("HX-Redirect", "/dashboard")
	log.Printf("Login successful for %s \n", username)
	return

}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	session, err := services.GetSession(r, "session")
	if err != nil {
		w.Header().Add("HX-Redirect", "/")
		return
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		log.Fatalf("Unable to save session: %s", err.Error())
	}

	w.Header().Add("HX-Redirect", "/")
}
