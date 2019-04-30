package apiserver

import (
	"net/http"
)

// Route is ..
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is ..
type Routes []Route

var routes = Routes{
	Route{
		"Status",
		"GET",
		"/status",
		Status,
	},
	Route{
		"UserIndex",
		"GET",
		"/users",
		UserIndex,
	},
	Route{
		"UserShow",
		"GET",
		"/users/{userId}",
		UserShow,
	},
	Route{
		"UserCreate",
		"POST",
		"/users",
		UserCreate,
	},
}
