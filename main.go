package main

import (
	"encoding/json"
	"fmt"
	"github.com/Mohannad-Zeido/HackerNewsScraper/html"
)

func main() {
	posts := html.GetPosts(2)

	pos, err := json.MarshalIndent(posts, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pos))

}
