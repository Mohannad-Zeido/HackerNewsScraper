package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"math"
	"strconv"
	"strings"
)

type responseItem struct {
	title    string
	uri      string
	author   string
	points   int
	comments int
	rank     int
}

func main() {
	c := colly.NewCollector()
	PostCountWanted := 2
	cou := float64(PostCountWanted) / float64(30)
	count := math.Ceil(cou)
	requestsDone := 0

	c.OnRequest(func(r *colly.Request) {
		requestsDone += 1
		fmt.Println("Visiting", r.URL.String())
	})

	var title []string
	var uri []string
	var author []string
	var point []int
	var numComments []int
	var rank []int

	c.OnHTML("body", func(e *colly.HTMLElement) {

		e.ForEachWithBreak("td > a.storylink", func(i int, element *colly.HTMLElement) bool {
			if len(title) >= PostCountWanted {
				return false
			}
			title = append(title, element.Text)
			uri = append(uri, element.Attr("href"))

			return true
		})

		e.ForEachWithBreak("td.subtext > a.hnuser", func(i int, element *colly.HTMLElement) bool {
			if len(author) >= PostCountWanted {
				return false
			}
			author = append(author, element.Text)
			return true
		})

		e.ForEachWithBreak("td.subtext > span.score", func(i int, element *colly.HTMLElement) bool {
			if len(point) >= PostCountWanted {
				return false
			}
			sco, _ := strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(element.Text, "points", "")))
			point = append(point, sco)
			return true
		})

		e.ForEachWithBreak("td.subtext > a[href^='item?id=']", func(i int, element *colly.HTMLElement) bool {
			if len(numComments) >= PostCountWanted {
				return false
			}
			com, _ := strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(element.Text, "comments", "")))
			numComments = append(numComments, com)
			return true
		})

		e.ForEachWithBreak("td > span.rank", func(i int, element *colly.HTMLElement) bool {
			if len(rank) >= PostCountWanted {
				return false
			}
			r, _ := strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(element.Text, ".", "")))
			rank = append(rank, r)
			return true
		})
	})

	c.OnScraped(func(response *colly.Response) {

		if requestsDone < int(count) {
			c.Visit("https://news.ycombinator.com/news?p=" + strconv.Itoa(requestsDone+1))
			return
		}
		fmt.Println("title")
		fmt.Println(title)
		fmt.Println("")
		fmt.Println("uri")
		fmt.Println(uri)
		fmt.Println("")
		fmt.Println("author")
		fmt.Println(author)
		fmt.Println("")
		fmt.Println("point")
		fmt.Println(point)
		fmt.Println("")
		fmt.Println("numComments")
		fmt.Println(numComments)
		fmt.Println("")
		fmt.Println("rank")
		fmt.Println(rank)
		fmt.Println("")
	})

	c.Visit("https://news.ycombinator.com")
}
