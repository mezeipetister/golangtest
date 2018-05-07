package main

import (
	"fmt"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/golang/glog"
	"log"
)

func demoController(w http.ResponseWriter, r *http.Request) {
	body := "Hello World"
	glog.Info("Hello")
	fmt.Fprint(w,body)
}

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", demoController).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
	glog.Flush()
}