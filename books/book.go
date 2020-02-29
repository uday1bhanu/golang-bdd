package books

import (
	"encoding/json"
	"strings"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b *Book) AuthorLastName() string {
	return strings.Fields(b.Author)[1]
}

func (b *Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	} else {
		return "SHORT STORY"
	}
}

func NewBookFromJSON(book []byte) (Book, error) {
	var r Book

	err := json.Unmarshal(book, &r)
	return r, err
}
