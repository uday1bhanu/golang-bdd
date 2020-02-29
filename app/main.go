package main

import (
	"log"

	b "github.com/uday1bhanu/golang-bdd/books"
)

func main() {
	longBook := b.Book{
		Title:  "Les Miserables",
		Author: "Victor Hugo",
		Pages:  1488,
	}

	shortBook := b.Book{
		Title:  "Fox In Socks",
		Author: "Dr. Seuss",
		Pages:  24,
	}

	log.Println(longBook.CategoryByLength())
	log.Println(shortBook.CategoryByLength())
}
