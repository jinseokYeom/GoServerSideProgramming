package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

// templates to be parsed
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// validate paths for security
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// define a page
type Page struct {
	Title string
	Body  []byte
}

// constructor
func NewPage(t string, b []byte) *Page {
	return &Page{Title: t, Body: b}
}

// save page to a text file
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// load from a text file
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return NewPage(title, body), nil
}

// render HTML template
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// wrapper function for handlers
func makeHandler(fn func(http.ResponseWriter,
	*http.Request, string)) http.HandlerFunc {
	// return a http.HandlerFunc type function
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, m[2])
	}
}

// view handler
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		// if page is not found, redirect to edit and create new
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

// edit handler
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = NewPage(title, nil)
	}

	renderTemplate(w, "edit", p)
}

// save handler
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := NewPage(title, []byte(body))

	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	fmt.Println("Server running...")

	http.ListenAndServe(":8080", nil)
}
