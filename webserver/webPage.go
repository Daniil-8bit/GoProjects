package webserver

import (
	"log"
	"net/http"
)

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	mainInfo := "Main page!"

	_, err := w.Write([]byte(mainInfo))
	if err != nil {
		log.Fatal(err)
	}
}

func consultingPageHandler(w http.ResponseWriter, r *http.Request) {
	consultingInfo := "Consulting in our company info page!"

	_, err := w.Write([]byte(consultingInfo))
	if err != nil {
		log.Fatal(err)
	}
}

func ActivateServer() {
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/consulting", consultingPageHandler)

	err := http.ListenAndServe("localhost:5555", nil)
	if err != nil {
		log.Fatal(err)
	}
}
