package main

import (
    "fmt"
    "model"
    "net/http"
)

func main() {
    http.HandleFunc("/view/", model.MakeHandler(model.ViewHandler))
    http.HandleFunc("/edit/", model.MakeHandler(model.EditHandler))
    http.HandleFunc("/save/", model.MakeHandler(model.SaveHandler))

    fmt.Println("Server running...")

    http.ListenAndServe(":8080", nil)
}
