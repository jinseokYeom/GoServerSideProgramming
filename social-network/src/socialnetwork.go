package main

import (
	"log"
	"net/http"
	"time"
)

const (
	INDEX   = "tmpl/index.html"
	SIGN_UP = "tmpl/signup.html"
	LOGIN   = "tmpl/login.html"
)

var (
	templates = template.Must(template.ParseFile(INDEX, SIGN_UP, LOGIN))
)

// define a page
type Page struct {
	Header  string
	Content []byte
}

// create a new page
func NewPage(h string, c []byte) *Page {
	return &Page{Header: h, Content: c}
}

// render templates
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// index page handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {

}

// sign up page handler
func SignUpHandler(w http.ResponseWriter, r *http.Request) {

}

// log in page handler
func LoginHandler(w http.ResponseWirter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/signup", SignupHandler)
	http.HandleFunc("/login", LoginHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
