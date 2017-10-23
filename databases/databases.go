package databases

import (
	"redditClone/indexers"
	"redditClone/models"
	"sync"
)

type Id int

type Database struct {
	topicsTable map[int]models.Topic
	indexer     *indexers.MinMaxHeap
	mux         sync.Mutex
}

func NewDatabase(querySize int) *Database {
	db := new(Database)
	db.topicsTable = make(map[int]models.Topic)
	db.indexer = indexers.NewMinMaxHeap(indexers.QueryMax, querySize)
	return db
}

func (db *Database) InsertTopic(topic models.Topic) error {
	db.mux.Lock()
	id := len(db.topicsTable)
	db.topicsTable[id] = topic
	db.indexer.AddById(id, topic.TotalVoteCount())
	db.mux.Unlock()
	return nil
}

func (db *Database) AddUpvote(id int) error {
	db.mux.Lock()
	db.topicsTable[id].ReceiveUpvote(1)
	db.indexer.Update(id, db.topicsTable[id].TotalVoteCount())
	db.mux.Unlock()
	return nil
}
