package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("Hello world"))
	})
	http.HandleFunc("/bye", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("bye world"))
	})
	http.ListenAndServe(":5000", nil)
}
