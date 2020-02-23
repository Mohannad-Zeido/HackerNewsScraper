package test

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/post"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPosts(t *testing.T) {
	posts, err := post.GetPosts(1, types.RunState{
		TestFile:  "testData/page",
		TestState: true,
	})
	assert.Nil(t, err)
	assert.Equal(t, "post1", posts[0].Title)
	assert.Equal(t, "https://www.google.com", posts[0].Uri)
	assert.Equal(t, 1, posts[0].Rank)

}
