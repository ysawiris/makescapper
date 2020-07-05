package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#cmn_wrap > div.content.js-table-sorting.floatfix.song-content > div > div > div.lfd-content.floatfix > div:nth-child(2)", func(e *colly.HTMLElement) {
		e.ForEach("div.lf-list__cell.lf-list__title.lf-list__cell--full > a[href]", func(_ int, e *colly.HTMLElement) {
			link := e.Attr("href")
			// Print link
			fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.lyricsfreak.com/d/drake/")
}
