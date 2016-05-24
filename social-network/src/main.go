package main

import (
	"controller"
	"log"
)

func main() {
	http.HandleFunc("/", controller.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
