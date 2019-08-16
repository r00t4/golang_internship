package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FibNumber struct {
	Current int `json:"current"`
	Prev    int `json:"prev"`
	Next    int `json:"next"`
}

func GetFib(i int) (FibNumber, error) {
	if i < 1 {
		return FibNumber{}, errors.New("number is negative")
	}
	return FibNumber{Current: getFib(i), Prev: getFib(i - 1), Next: getFib(i + 1)}, nil
}

func getFib(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 || i == 2 {
		return 1
	}
	return getFib(i-1) + getFib(i-2)
}

func handler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	keys := params["id"]

	key, err := strconv.Atoi(keys)
	if err != nil {
		fmt.Println(err)
		return
	}

	fib, err := GetFib(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(fib)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(data)
	log.Println(r.URL, fib.Current, r.Method)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/fibonacci/{id:[0-9]+}", handler).Methods("GET")
	http.Handle("/", rtr)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
