package models

type Topic struct {
	name     string `json: name`
	upvote   int32  `json: upvote`
	downvote int32  `json: downvote`
}
