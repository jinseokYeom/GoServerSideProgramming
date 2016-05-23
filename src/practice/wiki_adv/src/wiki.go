package main

import (
    "fmt"
    "server"
    "net/http"
)

func main() {
    http.HandleFunc("/view/", server.MakeHandler(server.ViewHandler))
    http.HandleFunc("/edit/", server.MakeHandler(server.EditHandler))
    http.HandleFunc("/save/", server.MakeHandler(server.SaveHandler))

    fmt.Println("Server running...")

    http.ListenAndServe(":8080", nil)
}
