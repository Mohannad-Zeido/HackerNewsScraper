package test

import (
	"fmt"
	"github.com/Mohannad-Zeido/HackerNewsScraper/post"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPosts_1Post(t *testing.T) {
	posts, err := post.GetPosts(1, types.RunState{TestState: true})
	assert.Nil(t, err)
	fmt.Println(posts)

}
