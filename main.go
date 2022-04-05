package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var countWithId map[int]int

var count int = 0

func main() {

	countWithId[1] = 10
	countWithId[2] = 20

	router := http.NewServeMux()
	router.HandleFunc("/increment", IncrementCount)
	router.HandleFunc("/decrement", DecrementCount)
	router.HandleFunc("/current", showCurrentCountWithId)

	log.Println("Starting server on :4001")
	err := http.ListenAndServe(":4001", router)
	if err != nil {
		log.Fatal(err)
	}
}

func IncrementCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	countEnteredByUser, _ := strconv.Atoi(r.URL.Query().Get("with"))

	if countEnteredByUser <= 0 {
		http.Error(w, "Please enter positive number", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "count was: %d\n", count)
	count = count + countEnteredByUser
	fmt.Fprintf(w, "count is : %d", count)
}

func DecrementCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	countEnteredByUser, _ := strconv.Atoi(r.URL.Query().Get("with"))
	fmt.Fprintf(w, "count was: %d\n", count)
	count = count - countEnteredByUser
	fmt.Fprintf(w, "count is : %d", count)
}

func showCurrentCountWithId(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("userId"))
	if id <= 0 {
		http.Error(w, "Please enter valid id ", http.StatusBadRequest)
		return
	}

	count, ok := countWithId[id]

	if !ok {
		http.Error(w, "User was not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "user %d count is :  %d", id, count)
}
