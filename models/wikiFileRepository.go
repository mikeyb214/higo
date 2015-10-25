//Package models contains data models for use in the wiki app.
package models

import "io/ioutil"

//WikiFileRepository is resposible for storage and retrieval of Wiki content to disk.
type WikiFileRepository struct {
}

const wikiPath = "./data/"
const wikiExtension = ".txt"

//SaveWiki stores the wiki content to disk.
func (r *WikiFileRepository) SaveWiki(w *Wiki) error {
	filename := wikiPath + w.Title + wikiExtension
	return ioutil.WriteFile(filename, w.Body, 0600)
}

//LoadWiki retrieves wiki content from disk.
func (r *WikiFileRepository) LoadWiki(title string) (*Wiki, error) {
	body, err := ioutil.ReadFile(wikiPath + title + wikiExtension)
	if err != nil {
		return nil, err
	}
	return &Wiki{Title: title, Body: body}, nil
}

//ListWikiTitles retrieves a list of wikis in the repository.
func (r *WikiFileRepository) ListWikiTitles() ([]string, error) {
	fileInfos, err := ioutil.ReadDir(wikiPath)
	if err != nil {
		return nil, err
	}
	wikiTitles := make([]string, len(fileInfos))
	for i, fileInfo := range fileInfos {
		wikiTitles[i] = fileInfo.Name()
	}
	return wikiTitles, nil
}
