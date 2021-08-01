package handler

import "net/http"

type AuthorizedHandler func(http.ResponseWriter, *http.Request)

func authorize(handlerFunc AuthorizedHandler, requiredPermissions ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		err := s.Authorize().AuthorizeToken(bearerToken, requiredPermissions)
		if err != nil {
			respondMsg(w, "Error: Could not authorize", http.StatusUnauthorized)
			return
		}

		handlerFunc(w, r)
	}
}
