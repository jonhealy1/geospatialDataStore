package main

import (
	"fmt"
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello from the web server")
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	log.Fatal(http.ListenAndServe(":12345", router))
}

/* Tutorial */
// https://www.thepolyglotdeveloper.com/2017/03/bundle-html-css-javascript-served-golang-application/
