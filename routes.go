package main

import (
	"net/http"

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
		Index,
	},
	Route{
		"SubmitTopic",
		"POST",
		"/submit",
		SubmitTopic,
	},
	Route{
		"Upvote",
		"POST",
		"/upvote/{topicId}",
		Upvote,
	},
	Route{
		"Downvote",
		"POST",
		"/downvote/{topicId}",
		Downvote,
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
