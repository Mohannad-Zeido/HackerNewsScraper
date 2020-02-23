package main

import (
	"encoding/json"
	"fmt"
	posts2 "github.com/Mohannad-Zeido/HackerNewsScraper/posts"
)

func main() {
	posts := posts2.GetPosts(35)
	fmt.Print("posts gotten ")
	fmt.Println(len(posts))
	pos, err := json.MarshalIndent(posts, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pos))

}
