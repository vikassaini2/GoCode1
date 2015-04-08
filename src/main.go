package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	//router.PathPrefix("/static/scripts/lib/").Handler(http.FileServer(http.Dir("./static/scripts/lib/")))

	//router.Handle("/js/{rest}", http.StripPrefix("/js/", http.FileServer(http.Dir(HomeFolder+"/jquery.js"))))

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8088", router))
}
