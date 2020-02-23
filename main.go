package main

import (
	"encoding/json"
	"fmt"
	"github.com/Mohannad-Zeido/HackerNewsScraper/post"
)

func main() {
	posts, err := post.GetPosts(35)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("number of post gotten ")
	fmt.Println(len(posts))
	pos, err := json.MarshalIndent(posts, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(pos))

}
