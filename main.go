package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"sync"
	"time"
)

type RequestPayload struct {
	ToSort [][]int `json:"to_sort"`
}

type ResponsePayload struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNs       int64   `json:"time_ns"`
}

func main() {
	http.HandleFunc("/process-single", processSingle)
	http.HandleFunc("/process-concurrent", processConcurrent)

	port := ":8000"
	http.ListenAndServe(port, nil)
}

func processSingle(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, false)
}

func processConcurrent(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, true)
}

func handleRequest(w http.ResponseWriter, r *http.Request, concurrent bool) {
	var requestPayload RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	startTime := time.Now()
	var sortedArrays [][]int

	if concurrent {
		sortedArrays = sortConcurrently(requestPayload.ToSort)
	} else {
		sortedArrays = sortSequentially(requestPayload.ToSort)
	}

	responsePayload := ResponsePayload{
		SortedArrays: sortedArrays,
		TimeNs:       time.Since(startTime).Nanoseconds(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responsePayload)
}

func sortSequentially(arrays [][]int) [][]int {
	for i := range arrays {
		sort.Ints(arrays[i])
	}
	return arrays
}

func sortConcurrently(arrays [][]int) [][]int {
	var wg sync.WaitGroup
	wg.Add(len(arrays))

	for i := range arrays {
		go func(i int) {
			defer wg.Done()
			sort.Ints(arrays[i])
		}(i)
	}

	wg.Wait()
	return arrays
}
