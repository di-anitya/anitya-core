package handlerv1

import (
	"context"
	"net/http"
	"os"
	"strings"

	common "../common"
	"../models"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtAuthentication handler
var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuthPost := []string{"/api/v1/project", "/api/v1/role", "/api/v1/user", "/api/v1/auth/tokens"} // List of endpoints that doesn't require auth
		requestPath := r.URL.Path                                                                         // current request path

		// check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuthPost {
			if value == requestPath && r.Method == "POST" {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") // Grab the token from the header

		if tokenHeader == "" { // Token is missing, returns with error code 403 Unauthorized
			response = common.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			common.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = common.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			common.Respond(w, response)
			return
		}

		tokenPart := splitted[1] // Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { // Malformed token, returns with http code 403 as usual
			response = common.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			common.Respond(w, response)
			return
		}

		if !token.Valid { // Token is invalid, maybe not signed on this server
			response = common.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			common.Respond(w, response)
			return
		}

		// Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		// fmt.Sprintf("User %", tk.UserID) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) // proceed in the middleware chain!
	})
}
