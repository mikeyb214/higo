//Package models contains data models for use in the wiki app.
package models

//Page is a simple representation of a web page.
type Page struct {
	PageTitle  string
	WikiTitles []string
	Wiki       *Wiki
}
