package main

import (
	"AEScrypto"
	"fmt"
	"io/ioutil"
	"log"
	"model"
	"net/http"
)

// generate an universal key
func generateKey() error {
	key, err := AEScrypto.RandomKey()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data/key", key, 0600)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := generateKey()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/view/", model.MakeHandler(model.ViewHandler))
	http.HandleFunc("/edit/", model.MakeHandler(model.EditHandler))
	http.HandleFunc("/save/", model.MakeHandler(model.SaveHandler))

	fmt.Println("Server running...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
