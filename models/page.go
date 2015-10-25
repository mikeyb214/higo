//Package models implements a page and storage mechanism
package models

import "io/ioutil"

//Page is a simple representation of a web page.
type Page struct {
	Title string
	Body  []byte
}

//SavePage stores the page to disk.
func (p *Page) SavePage() error {
	filename := "./data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//LoadPage retrievs the page from disk.
func LoadPage(title string) (*Page, error) {
	filename := "./data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
