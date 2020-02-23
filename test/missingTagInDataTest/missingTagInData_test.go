package missingTagInDataTest

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/post"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissingTagInDataTest_HTMLTagOnly(t *testing.T) {
	numberOfPosts := 1
	posts, err := post.GetPosts(numberOfPosts, types.RunState{
		TestFile:  "testData/emptyPage",
		TestState: true,
	})
	assert.EqualError(t, err, types.ErrGettingPostsTableNode)
	assert.Nil(t, posts)
}

func TestMissingTagInDataTest_NoTbody(t *testing.T) {
	numberOfPosts := 1
	posts, err := post.GetPosts(numberOfPosts, types.RunState{
		TestFile:  "testData/noTBody",
		TestState: true,
	})
	assert.EqualError(t, err, types.ErrGettingPostsTbodyNode)
	assert.Nil(t, posts)
}

func TestMissingTagInDataTest_NoPosts(t *testing.T) {
	numberOfPosts := 1
	posts, err := post.GetPosts(numberOfPosts, types.RunState{
		TestFile:  "testData/noPosts",
		TestState: true,
	})
	assert.EqualError(t, err, types.ErrNoPostsOnPage)
	assert.Nil(t, posts)
}

func TestMissingTagInDataTest_PostNot3Rows(t *testing.T) {
	numberOfPosts := 2
	posts, err := post.GetPosts(numberOfPosts, types.RunState{
		TestFile:  "testData/postIncomplete",
		TestState: true,
	})
	assert.EqualError(t, err, types.ErrGettingNextPost)
	assert.Nil(t, posts)
}
