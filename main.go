package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/increment", IncrementCount)
	router.HandleFunc("/decrement", DecrementCount)
	router.HandleFunc("/current", showCurrentCount)

	log.Println("Starting server on :4001")
	err := http.ListenAndServe(":4001", router)
	log.Fatal(err)
}

type Count int

var count Count = 0

func IncrementCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method is not allowed", 405)
		return 
	}

	countEnteredByUser, _ := strconv.Atoi(r.URL.Query().Get("with"))
	
	if countEnteredByUser <= 0 {
		http.Error(w, "Please enter positive number", 400)
		return
	}

	fmt.Fprintf(w, "count was: %d\n", count)
	count = count + Count(countEnteredByUser)
	fmt.Fprintf(w, "count is : %d", count)
}

func DecrementCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method is not allowed", 405)
		return
	}
	countEnteredByUser, _ := strconv.Atoi(r.URL.Query().Get("with"))
	fmt.Fprintf(w, "count was: %d\n", count)
	count = count - Count(countEnteredByUser)
	fmt.Fprintf(w, "count is : %d", count)
}

func showCurrentCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, "method is not allowed", 405)
		return 
	}
	
	fmt.Fprintf(w, "current count was: %d", count)
}
