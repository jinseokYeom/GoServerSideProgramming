package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	LOGIN = "tmpl/login.html"
	MLIST = "tmpl/members.html"
)

var tmpls = template.Must(template.ParseFiles(LOGIN, MLIST))

// define a user
type User struct {
	Id        string // user ID
	Password  []byte // encrypted password
	UserName  string // user name
	FirstName string // first name
	LastName  string // last name
	Gender    string // gender
}

// create a new user
func NewUser() *User {
	return &User{}
}

// register a new user
func Register() (*User, error) {

	// to be implemented
	return NewUser(), nil
}

// existing user login
func (u *User) Login() {

}

// look up user from user data
func userLookUp(username string) error {

	return nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	log.Printf("Request method: %s\n", m)

	if m == "GET" {
		err := tmpls.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		r.ParseForm()
		username := r.FormValue("username")
		err := userLookUp(username)
		if err != nil {
			// redirect to register page
			//return
		}

		log.Printf("\033[1m%s\033[0m logged in\n", username)
		// redirect to member page
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)

	log.Println("Server running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
