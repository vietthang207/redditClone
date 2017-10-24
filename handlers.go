package main

import (
	"html/template"
	"log"
	"net/http"
	"redditClone/models"
	"strconv"
	"strings"
)

type TopicData struct {
	Rank          int
	Id            int
	Name          string
	UpvoteCount   int
	DownvoteCount int
}

func convertData(queryResult []models.Topic) []TopicData {
	ret := make([]TopicData, len(queryResult))
	for i := 0; i < len(queryResult); i++ {
		ret[i] = TopicData{i + 1, queryResult[i].Id(), queryResult[i].Name(), queryResult[i].UpvoteCount(), queryResult[i].DownvoteCount()}
	}
	return ret
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle index ")
	mock := convertData(db.GetTopTopics())
	t := template.New("home.html")
	t, err := t.ParseFiles("templates/home.html")
	log.Println(err)
	t.Execute(w, mock)
}

func SubmitTopic(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle submit ")
	r.ParseForm()
	formData := r.PostFormValue("submitText")
	log.Println("Create topic:", formData)
	db.InsertTopic(formData)
	log.Println("Topics count: ", db.Size())
}

func Upvote(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle upvote ")
	tokens := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(tokens[len(tokens)-1])
	log.Println("Upvote ", id)
	db.AddUpvote(id)
}

func Downvote(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle downvote ")
	tokens := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(tokens[len(tokens)-1])
	log.Println("Downvote ", id)
	db.AddDownvote(id)
}
