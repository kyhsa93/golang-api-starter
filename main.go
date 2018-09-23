package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprintln(writer, "routing index")
}

func hello(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprintln(writer, "routing hello")
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello", hello)
	log.Fatal(http.ListenAndServe(":5000", router))
}
