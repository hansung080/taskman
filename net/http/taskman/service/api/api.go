package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hansung080/taskman/net/http/taskman/service"
	"github.com/hansung080/taskman/task"
)

const PathPrefix = "/api/v1/tasks/"

func getTasks(r *http.Request) ([]task.Task, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	encodedTasks, exists := r.PostForm["task"]
	if !exists {
		return nil, errors.New("task not provided")
	}

	var tasks []task.Task
	for _, encodedTask := range encodedTasks {
		var t task.Task
		if err := json.Unmarshal([]byte(encodedTask), &t); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func Post(w http.ResponseWriter, r *http.Request) {
	tasks, err := getTasks(r)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%s %s\n", http.MethodPost, PathPrefix)
	log.Println(tasks)

	for _, t := range tasks {
		id, err := service.TaskAccessor.Add(t)
		if err = json.NewEncoder(w).Encode(service.Response{
			ID:    id,
			Task:  t,
			Error: service.ResponseError{Err: err},
		}); err != nil {
			log.Println(err)
			return
		}
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	log.Printf("%s %s%s\n", http.MethodGet, PathPrefix, id)

	t, err := service.TaskAccessor.Get(id)
	if err = json.NewEncoder(w).Encode(service.Response{
		ID:    id,
		Task:  t,
		Error: service.ResponseError{Err: err},
	}); err != nil {
		log.Println(err)
		return
	}
}

func Put(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	tasks, err := getTasks(r)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%s %s%s\n", http.MethodPut, PathPrefix, id)
	log.Println(tasks)

	for _, t := range tasks {
		err = service.TaskAccessor.Update(id, t)
		if err = json.NewEncoder(w).Encode(service.Response{
			ID:    id,
			Task:  t,
			Error: service.ResponseError{Err: err},
		}); err != nil {
			log.Println(err)
			return
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	log.Printf("%s %s%s\n", http.MethodDelete, PathPrefix, id)

	err := service.TaskAccessor.Delete(id)
	if err = json.NewEncoder(w).Encode(service.Response{
		ID:    id,
		Error: service.ResponseError{Err: err},
	}); err != nil {
		log.Println(err)
		return
	}
}
