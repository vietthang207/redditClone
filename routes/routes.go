package routes

import (
	"net/http"
	"redditClone/handlers"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"SubmitTopic",
		"POST",
		"/submit",
		handlers.SubmitTopic,
	},
	Route{
		"Upvote",
		"POST",
		"/upvote",
		handlers.Upvote,
	},
	Route{
		"Downvote",
		"POST",
		"/downvote",
		handlers.Downvote,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
