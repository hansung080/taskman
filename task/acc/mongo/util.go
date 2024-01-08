package mongo

import (
	"github.com/globalsign/mgo/bson"
	"github.com/hansung080/taskman/task"
)

func toObjectId(id task.ID) bson.ObjectId {
	return bson.ObjectIdHex(string(id))
}

func fromObjectId(objID bson.ObjectId) task.ID {
	return task.ID(objID.Hex())
}
