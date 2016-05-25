package main

import ()

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

}

// existing user login
func (u *User) Login() {

}

func renderTemplate() {

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	u, err := Register()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "register")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
