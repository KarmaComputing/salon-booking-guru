package handler

import "net/http"

type AuthorizedHandler func(http.ResponseWriter, *http.Request)

// Authorizes a request based on the "Authorization" header, and ensures the
// token has access to the specified requiredPermissions.
//
// Returns a HandlerFunc to be consumed by the mux router.
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
