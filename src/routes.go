package main

import (
	"github.com/gorilla/mux"
	"net/http"
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
	HomeFolder := "./static/"
	HomeFolderLib := "./static/scripts/lib"
	router.PathPrefix("/lib/").Handler(http.StripPrefix("/lib/", http.FileServer(http.Dir(HomeFolderLib+"/"))))
	router.PathPrefix("/views/").Handler(http.StripPrefix("/views/", http.FileServer(http.Dir(HomeFolder+"/"))))

	for _, route := range routes {
		// var handler http.Handler

		//handler = route.HandlerFunc
		//  handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}

	return router
}

var routes = Routes{

	Route{
		"Login",
		"POST",
		"/login",
		Login,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
}
