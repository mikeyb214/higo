//Package models contains data models for use in the wiki app.
package models

//WikiRepository is an interface for storage and retrieval of Wiki's.
type WikiRepository interface {
	SaveWiki(wiki *Wiki) (err error)
	LoadWiki(title string) (wiki *Wiki, err error)
	ListWikiTitles() (wikiTitles []string, err error)
}
