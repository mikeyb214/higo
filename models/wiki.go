//Package models contains data models for use in the wiki app.
package models

//Wiki is a simple representation of a wiki article.
type Wiki struct {
	Title string
	Body  []byte
}
