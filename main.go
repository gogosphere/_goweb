package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/describe-instances", JsonFull)
	router.HandleFunc("/{calledfile}", CalledFile)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	file, e := ioutil.ReadFile("data/index.html")
	if e != nil {
		fmt.Fprintln(w, e, "ioutil read bork")
	}
	/*
		_, err := json.Marshal(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintln(w, string(file))
}

func JsonFull(w http.ResponseWriter, r *http.Request) {

	file, e := ioutil.ReadFile("describe-instances.json")
	if e != nil {
		fmt.Fprintln(w, e, "ioutil read bork")
	}
	/*
		_, err := json.Marshal(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintln(w, string(file))
}

func CalledFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	calledfile := vars["calledfile"]
	file, e := ioutil.ReadFile("data/" + calledfile)
	if e != nil {
		fmt.Fprintln(w, e, "ioutil read bork")
	}
	/*
		_, err := json.Marshal(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintln(w, string(file))
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
