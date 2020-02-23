package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Mohannad-Zeido/HackerNewsScraper/post"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
)

func main() {

	var numberOfPosts int

	flag.IntVar(&numberOfPosts, "posts", -1, "how many posts to print. A positive integer <= 100")
	flag.IntVar(&numberOfPosts, "p", -1, "how many posts to print. A positive integer <= 100  (shorthand)")

	flag.Parse()
	if numberOfPosts < 0 || numberOfPosts > 100 {
		fmt.Println("please input a valid number of posts to get (A positive integer <= 100)")
		return
	}
	posts, err := post.GetPosts(numberOfPosts, types.RunState{TestState: false})
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
