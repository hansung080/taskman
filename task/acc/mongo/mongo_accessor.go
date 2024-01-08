package mongo

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/hansung080/taskman/task"
)

type MongoAccessor struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func NewMongoAccessor(url, db, collection string) *MongoAccessor {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Panic(err)
	}
	c := session.DB(db).C(collection)
	return &MongoAccessor{
		session:    session,
		collection: c,
	}
}

func (m *MongoAccessor) Close() error {
	m.session.Close()
	return nil
}

func (m *MongoAccessor) Add(t task.Task) (task.ID, error) {
	objID := bson.NewObjectId()
	_, err := m.collection.UpsertId(objID, t)
	return fromObjectId(objID), err
}

func (m *MongoAccessor) Update(id task.ID, t task.Task) error {
	return m.collection.UpdateId(toObjectId(id), t)
}

func (m *MongoAccessor) Delete(id task.ID) error {
	return m.collection.RemoveId(toObjectId(id))
}

func (m *MongoAccessor) Get(id task.ID) (task.Task, error) {
	t := task.Task{}
	err := m.collection.FindId(toObjectId(id)).One(&t)
	return t, err
}
