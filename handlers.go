package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func SubmitTopic(w http.ResponseWriter, r *http.Request) {
}

func Upvote(w http.ResponseWriter, r *http.Request) {
}

func Downvote(w http.ResponseWriter, r *http.Request) {
}
