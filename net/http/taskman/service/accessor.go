package service

import (
	"github.com/hansung080/taskman/task"
	"github.com/hansung080/taskman/task/acc/mongo"
)

// FIXME: TaskAccessor is not thread-safe.
// var TaskAccessor task.Accessor = mem.NewMemoryAccessor()
// var TaskAccessor task.Accessor = mongo.NewMongoAccessor("mongodb://localhost:27017/test", "test", "tasks")
var TaskAccessor task.Accessor = mongo.NewMongoAccessor("localhost:27017/test", "test", "tasks")
