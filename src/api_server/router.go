package apiserver

import (
	"net/http"

	"../common"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// NewRouter is ..
func NewRouter() *mux.Router {
	mw := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	var router = mux.NewRouter()
	router.HandleFunc("/api/auth/token", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("no auth required\n"))
	}).Methods("GET")

	var _router = mux.NewRouter()

	an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(_router))
	router.PathPrefix("/api").Handler(an)

	var v1 = router.PathPrefix("/v1").Subrouter()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = common.Logger(handler, route.Name)

		v1.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return v1
}
