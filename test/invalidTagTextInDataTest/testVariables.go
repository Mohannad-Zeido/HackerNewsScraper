package invalidTagTextInDataTest

import "github.com/Mohannad-Zeido/HackerNewsScraper/types"

var (
	testDataPath     = "testData/page"
	AllExpectedPosts = []types.Post{
		{
			Title:    "post3",
			Uri:      "https://www.post3.com",
			Author:   "post3Author",
			Points:   128,
			Comments: 34,
			Rank:     3,
		},
		{
			Title:    "post4",
			Uri:      "https://www.post4.com",
			Author:   "post4Author",
			Points:   77,
			Comments: 33,
			Rank:     4,
		},
		{
			Title:    "post5",
			Uri:      "https://www.post5.com",
			Author:   "post5Author",
			Points:   39,
			Comments: 27,
			Rank:     5,
		},
		{
			Title:    "post7",
			Uri:      "https://www.post7.com",
			Author:   "post7Author",
			Points:   285,
			Comments: 125,
			Rank:     7,
		},
		{
			Title:    "post9",
			Uri:      "https://www.post9.com",
			Author:   "post9Author",
			Points:   11,
			Comments: 0,
			Rank:     9,
		},
		{
			Title:    "post10",
			Uri:      "https://www.post10.com",
			Author:   "post10Author",
			Points:   123,
			Comments: 16,
			Rank:     10,
		},
		{
			Title:    "post12",
			Uri:      "https://www.post12.com",
			Author:   "post12Author",
			Points:   91,
			Comments: 14,
			Rank:     12,
		},
		{
			Title:    "post13",
			Uri:      "https://www.post13.com",
			Author:   "post13Author",
			Points:   136,
			Comments: 58,
			Rank:     13,
		},
		{
			Title:    "post14",
			Uri:      "https://www.post14.com",
			Author:   "post14Author",
			Points:   204,
			Comments: 260,
			Rank:     14,
		},
		{
			Title:    "post16",
			Uri:      "https://www.post16.com",
			Author:   "post16Author",
			Points:   52,
			Comments: 21,
			Rank:     16,
		},
		{
			Title:    "post18",
			Uri:      "https://www.post18.com",
			Author:   "post18Author",
			Points:   14,
			Comments: 0,
			Rank:     18,
		},
	}
)
