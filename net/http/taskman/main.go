package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hansung080/taskman/net/http/taskman/service/api"
	"github.com/hansung080/taskman/net/http/taskman/service/html"
)

const idPattern = "/{id:[0-9]+}"

func main() {
	router := mux.NewRouter()
	router.PathPrefix(html.PathPrefix).
		Path(idPattern).
		Methods(http.MethodGet).
		HandlerFunc(html.Get)

	sub := router.PathPrefix(api.PathPrefix).Subrouter()
	sub.HandleFunc("/", api.Post).Methods(http.MethodPost)
	sub.HandleFunc(idPattern, api.Get).Methods(http.MethodGet)
	sub.HandleFunc(idPattern, api.Put).Methods(http.MethodPut)
	sub.HandleFunc(idPattern, api.Delete).Methods(http.MethodDelete)

	http.Handle("/", router)
	http.Handle(
		"/css/",
		http.StripPrefix(
			"/css/",
			http.FileServer(http.Dir("styles")),
		),
	)

	log.Println("TaskMan is running at 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
