package missingTagTextInDataTest

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/post"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPosts_MissingTagText_1Post(t *testing.T) {
	numberOfPosts := 1
	posts, err := post.GetPosts(numberOfPosts, types.RunState{
		TestFile:  testDataPath,
		TestState: true,
	})
	assert.Nil(t, err)
	assert.Equal(t, numberOfPosts, len(posts))
	checkPosts(AllExpectedPosts, posts, numberOfPosts, t)
}

func TestGetPosts_MissingTagText_LastPostOnNextPage(t *testing.T) {
	numberOfPosts := 6
	posts, err := post.GetPosts(numberOfPosts, types.RunState{
		TestFile:  testDataPath,
		TestState: true,
	})
	assert.Nil(t, err)
	assert.Equal(t, numberOfPosts, len(posts))
	checkPosts(AllExpectedPosts, posts, numberOfPosts, t)
}

func TestGetPosts_MissingTagText_WithInvalidPosts(t *testing.T) {
	numberOfPosts := 9
	posts, err := post.GetPosts(numberOfPosts, types.RunState{
		TestFile:  testDataPath,
		TestState: true,
	})
	assert.Nil(t, err)
	assert.Equal(t, numberOfPosts, len(posts))
	checkPosts(AllExpectedPosts, posts, numberOfPosts, t)
}

func TestGetPosts_MissingTagText_AllPosts(t *testing.T) {
	numberOfPosts := 12
	posts, err := post.GetPosts(numberOfPosts, types.RunState{
		TestFile:  testDataPath,
		TestState: true,
	})
	assert.Nil(t, err)
	assert.Equal(t, numberOfPosts, len(posts))
	checkPosts(AllExpectedPosts, posts, numberOfPosts, t)
}

func TestGetPosts_MissingTagText_notEnoughPostData(t *testing.T) {
	numberOfPosts := 13
	_, err := post.GetPosts(numberOfPosts, types.RunState{
		TestFile:  testDataPath,
		TestState: true,
	})
	assert.NotNil(t, err)
}

//todo this needs to go to a top level file
func checkPosts(expectedPosts, actualPosts []types.Post, expectedNumberOfPosts int, t *testing.T) {

	for i := 0; i < expectedNumberOfPosts; i++ {
		t.Logf("checking post %d", i+1)
		assert.Equal(t, expectedPosts[i].Title, actualPosts[i].Title)
		assert.Equal(t, expectedPosts[i].Uri, actualPosts[i].Uri)
		assert.Equal(t, expectedPosts[i].Rank, actualPosts[i].Rank)
		assert.Equal(t, expectedPosts[i].Points, actualPosts[i].Points)
		assert.Equal(t, expectedPosts[i].Author, actualPosts[i].Author)
		assert.Equal(t, expectedPosts[i].Comments, actualPosts[i].Comments)

	}
}
