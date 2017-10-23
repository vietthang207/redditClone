package databases

import (
	"redditClone/indexers"
	"redditClone/models"
	"sync"
)

type Id int

type Database struct {
	topicsTable map[int]*models.Topic
	indexer     *indexers.MinMaxHeap
	mux         sync.Mutex
}

func NewDatabase(querySize int) *Database {
	db := new(Database)
	db.topicsTable = make(map[int]*models.Topic)
	db.indexer = indexers.NewMinMaxHeap(indexers.QueryMax, querySize)
	return db
}

func (db *Database) Size() int {
	db.mux.Lock()
	ret := len(db.topicsTable)
	db.mux.Unlock()
	return ret
}

func (db *Database) GetTopTopics() []models.Topic {
	db.mux.Lock()
	tmp := db.indexer.Query()
	ret := make([]models.Topic, len(tmp))
	for i := 0; i < len(tmp); i++ {
		ret[i] = *db.topicsTable[tmp[i].Id()]
	}
	db.mux.Unlock()
	return ret
}

func (db *Database) InsertTopic(name string) error {
	db.mux.Lock()
	id := len(db.topicsTable)
	topic := models.NewTopic(name, id)
	db.topicsTable[id] = &topic
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

func (db *Database) AddDownvote(id int) error {
	db.mux.Lock()
	db.topicsTable[id].ReceiveDownvote(1)
	db.indexer.Update(id, db.topicsTable[id].TotalVoteCount())
	db.mux.Unlock()
	return nil
}
