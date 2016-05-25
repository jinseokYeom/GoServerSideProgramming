package main

import ()

// define a page
type Page struct {
	Header  string
	Content []byte
}

// create a new page
func NewPage(h string, c []byte) *Page {
	return &Page{Header: h, Content: c}
}

//
