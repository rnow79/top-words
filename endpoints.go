package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// Default top size.
var defaultTopSize int = 10

// Maximum top value.
var maximumTopValue int = 100

// Redirect to main page of the front end.
func redirectToFront(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/front/index.htm", http.StatusFound)
}

// TopWords: this endpoint will parse input text and return the most used words.
func postTopWords(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Load default top.
	var top int = defaultTopSize
	var ignorecase bool
	if n := r.PostForm.Get("n"); len(n) > 0 {
		n, err := strconv.Atoi(n)
		if err != nil {
			sendBadRequest(w, "n must be numeric")
			return
		}
		if n < 1 || n > maximumTopValue {
			sendBadRequest(w, fmt.Sprintf("n must be between 1 and %d", maximumTopValue))
			return
		}
		// 0 < n <= maxtop, so load it as new top.
		top = n
	}
	if r.PostForm.Get("ignorecase") == "1" {
		ignorecase = true
	}
	// Get input text.
	text := r.PostForm.Get("text")
	if len(text) == 0 {
		sendBadRequest(w, "text cannot be empty")
		return
	}
	// Replace some chars for better performance.
	text = strings.ReplaceAll(text, ",", " ")
	text = strings.ReplaceAll(text, ". ", " ")
	text = strings.ReplaceAll(text, " .", " ")
	text = strings.ReplaceAll(text, "(", " ")
	text = strings.ReplaceAll(text, ")", " ")
	text = strings.ReplaceAll(text, "[", " ")
	text = strings.ReplaceAll(text, "]", " ")
	text = strings.ReplaceAll(text, "{", " ")
	text = strings.ReplaceAll(text, "}", " ")
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")
	text = strings.ReplaceAll(text, "|", " ")
	text = strings.ReplaceAll(text, "@", " ")
	text = strings.ReplaceAll(text, ":", " ")
	// Create a frequency map and insert all text words.
	frequency := make(map[string]int)
	for _, word := range strings.Split(text, " ") {
		if len(word) == 0 {
			continue
		}
		if ignorecase {
			word = strings.ToLower(word)
		}
		frequency[word]++
	}
	// Now sort. Since package sort has Slice method, we have to convert map to struct.
	// In addition, we can define here the final JSON output.
	type Result struct {
		Word  string `json:"word"`
		Count int    `json:"count"`
	}
	var results []Result
	for k, v := range frequency {
		results = append(results, Result{k, v})
	}
	// Now we can sort.
	sort.Slice(results, func(current, next int) bool {
		return results[current].Count > results[next].Count
	})
	// Subslice results only if its length > top
	if len(results) > top {
		results = results[:top]
	}
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	json.NewEncoder(w).Encode(results)
}

// Send errors as 400 - Bad Request.
func sendBadRequest(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusBadRequest)
}
