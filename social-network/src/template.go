package main

import ()

const (
	INDEX   = "tmpl/index.html"
	SIGN_UP = "tmpl/sign_up.html"
	LOGIN   = "tmpl/login.html"
)

var templates = template.Must(template.ParseFiles(INDEX, SIGN_UP, LOGIN))

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		html.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
