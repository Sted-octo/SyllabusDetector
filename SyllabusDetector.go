package main

import (
	"SyllabusDetector/controllers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", controllers.Welcome)

	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
