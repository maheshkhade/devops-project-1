package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Hit")
	fmt.Print("hiiiiiii")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home)
	err := http.ListenAndServe(":8080", myRouter)
	if err != nil {
		log.Println(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home Hit")
	fmt.Fprintf(w, "This is a home page 25 Apr")
}
