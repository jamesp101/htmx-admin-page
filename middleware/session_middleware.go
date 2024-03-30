package middleware

import (
	"context"
	"main/services"
	"net/http"
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := services.GetSession(r, "session")
		if err != nil {
			w.Header().Add("HX-Redirect", "/l")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		username, ok := session.Values["username"]
		if !ok {
			w.Header().Add("HX-Redirect", "/")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		usernameStr, ok := username.(string)
		if !ok {
			w.Header().Add("HX-Redirect", "/")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", usernameStr)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
