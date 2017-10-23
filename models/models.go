package models

type Topic struct {
	id            int    `json: id`
	name          string `json: name`
	upvoteCount   int    `json: upvote`
	downvoteCount int    `json: downvote`
}

func NewTopic(name string, id int) Topic {
	return Topic{id, name, 0, 0}
}

func (t Topic) Id() int {
	return t.id
}

func (t Topic) Name() string {
	return t.name
}
func (t Topic) UpvoteCount() int {
	return t.upvoteCount
}

func (t Topic) ReceiveUpvote(n int) {
	t.upvoteCount += n
}

func (t Topic) ReceiveDownvote(n int) {
	t.downvoteCount += n
}

func (t Topic) DownvoteCount() int {
	return t.downvoteCount
}

func (t Topic) TotalVoteCount() int {
	return t.upvoteCount - t.downvoteCount
}
