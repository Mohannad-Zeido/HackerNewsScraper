package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)
func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// On every a element which has href attribute call callback
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		if e.Attr("class") == "athing"{
			fmt.Println(e)
		}
		//link := e.Attr("href")
		//// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		//// Visit link found on page
		//// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnHTML("td", func(e *colly.HTMLElement) {
		if e.Attr("class") == "subtext"{
			fmt.Println(e)
		}

		if e.Attr("class") == "morelink"{
			fmt.Println(e)
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "morelink"{
			fmt.Println(e.Attr("href"))
		}
	})

	c.Visit("https://news.ycombinator.com/")
}
