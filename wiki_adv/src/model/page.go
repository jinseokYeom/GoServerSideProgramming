package model

import (
    "regexp"
    "AEScrypto"
    "net/http"
    "io/ioutil"
    "html/template" 
)

const (
    EDIT    = "tmpl/edit.html"  // edit page html template
    VIEW    = "tmpl/view.html"  // view page html template
)

var (
    key, _      = AEScrypto.RandomKey()
    templates   = template.Must(template.ParseFiles(EDIT, VIEW))
    validPath   = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
)

// define a page
type Page struct {
    Title   string
    Body    []byte
}

// constructor
func NewPage(title string, body []byte) *Page {
    return &Page{Title: title, Body: body}
}

// load from a text file
func loadPage(title string) (*Page, error) {
    fileName := title + ".txt"

    body, err := ioutil.ReadFile("data/" + fileName)
    if err != nil {
        return nil, err
    }

    return NewPage(title, body), nil
}

// decrypt and load page
func decryptedLoadPage(key []byte, title string) (*Page, error) {
    fileName := title + ".txt"
    
    body, err := ioutil.ReadFile("data/" + fileName)
    if err != nil {
        return nil, err
    }

    // decrypt the content
    decrypted, err := AEScrypto.AESDecrypt(key, body)
    if err != nil {
        return nil, err
    }

    return NewPage(title, decrypted), nil
}

// save page to a text file
func (p *Page) save() error {
    fileName := p.Title + ".txt"
    return ioutil.WriteFile("data/" + fileName, p.Body, 0600)
}

// encrypt and save page to a text file
func (p *Page) encryptedSave(key []byte) error {
    fileName := p.Title + ".txt"
    
    // encrypt the content
    encryptedBody, err := AEScrypto.AESEncrypt(key, p.Body)
    if err != nil {
        return err
    }

    return ioutil.WriteFile("data/" + fileName, encryptedBody, 0600)
}

// render HTML template
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl + ".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// wrapper function for handlers
func MakeHandler(fn func (http.ResponseWriter,
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
func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := decryptedLoadPage(key, title)
    if err != nil {
        // if page is not found, redirect to edit and create new
        http.Redirect(w, r, "/edit/" + title, http.StatusFound)
        return
    }

    renderTemplate(w, "view", p)
}

// edit handler
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := decryptedLoadPage(key, title)
    if err != nil {
        p = NewPage(title, nil)
    }

    renderTemplate(w, "edit", p)
}

// save handler
func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := NewPage(title, []byte(body))
    
    err := p.encryptedSave(key)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/view/" + title, http.StatusFound)
}
