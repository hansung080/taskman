package html

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/hansung080/taskman/net/http/taskman/service"
	"github.com/hansung080/taskman/task"
)

const PathPrefix = "/tasks/"

var tmpl = template.Must(template.ParseGlob("views/*.html"))

func Get(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	log.Printf("%s %s%s\n", http.MethodGet, PathPrefix, id)

	t, err := service.TaskAccessor.Get(id)
	if err = tmpl.ExecuteTemplate(w, "task.html", &service.Response{
		ID:    id,
		Task:  t,
		Error: service.ResponseError{Err: err},
	}); err != nil {
		log.Println(err)
		return
	}
}
