package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = WebRequestWrapper(route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func WebRequestWrapper(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		Logger.Printf(
			"%v\t%v\t%v\t%v",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start).String(),
		)
	})
}

var routes = Routes{
	Route{
		"Version",
		"GET",
		"/",
		Version,
	},
	Route{
		"GetBoxes",
		"GET",
		"/Boxes",
		GetBoxes,
	},
	Route{
		"GetBox",
		"GET",
		"/GetBox/{name}",
		GetBox,
	},
	Route{
		"CreateBox",
		"POST",
		"/CreateBox",
		CreateBox,
	},
}
