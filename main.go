package main

import (
	"encoding/json"
	"fmt"
	"github.com/Mohannad-Zeido/HackerNewsScraper/post"
)

func main() {
	posts := post.GetPosts(35)
	fmt.Print("number of post gotten ")
	fmt.Println(len(posts))
	pos, err := json.MarshalIndent(posts, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pos))

}
