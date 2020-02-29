package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	b "github.com/uday1bhanu/golang-bdd/books"
	c "github.com/uday1bhanu/golang-bdd/cart"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

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

	log.Info().Msg(longBook.CategoryByLength())
	log.Info().Msg(shortBook.CategoryByLength())

	cart := c.Cart{}
	itemA := c.Item{ID: "itemA", Name: "Item A", Price: 10.20, Qty: 0}
	itemB := c.Item{ID: "itemB", Name: "Item B", Price: 7.66, Qty: 0}
	cart.AddItem(itemA)
	cart.AddItem(itemB)

	originalItemCount := cart.TotalUniqueItems()
	originalUnitCount := cart.TotalUnits()
	originalAmount := cart.TotalAmount()
	log.Info().Int("originalItemCount", originalItemCount).Int("originalUnitCount", originalUnitCount).Float64("originalAmount", originalAmount).Msg("")
}
