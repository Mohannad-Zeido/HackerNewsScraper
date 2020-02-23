package missingTagTextInDataTest

import "github.com/Mohannad-Zeido/HackerNewsScraper/types"

var (
	testDataPath     = "testData/page"
	AllExpectedPosts = []types.Post{
		{
			Title:    "post2",
			URI:      "https://www.post2.com",
			Author:   "post2Author",
			Points:   39,
			Comments: 3,
			Rank:     2,
		},
		{
			Title:    "post3",
			URI:      "https://www.post3.com",
			Author:   "post3Author",
			Points:   128,
			Comments: 34,
			Rank:     3,
		},
		{
			Title:    "post4",
			URI:      "https://www.post4.com",
			Author:   "post4Author",
			Points:   77,
			Comments: 33,
			Rank:     4,
		},
		{
			Title:    "post5",
			URI:      "https://www.post5.com",
			Author:   "post5Author",
			Points:   39,
			Comments: 27,
			Rank:     5,
		},
		{
			Title:    "post6",
			URI:      "https://www.post6.com",
			Author:   "post6Author",
			Points:   402,
			Comments: 230,
			Rank:     6,
		},
		{
			Title:    "post7",
			URI:      "https://www.post7.com",
			Author:   "post7Author",
			Points:   285,
			Comments: 125,
			Rank:     7,
		},
		{
			Title:    "post10",
			URI:      "https://www.post10.com",
			Author:   "post10Author",
			Points:   123,
			Comments: 16,
			Rank:     10,
		},
		{
			Title:    "post14",
			URI:      "https://www.post14.com",
			Author:   "post14Author",
			Points:   204,
			Comments: 260,
			Rank:     14,
		},
		{
			Title:    "post15",
			URI:      "https://www.post15.com",
			Author:   "post15Author",
			Points:   35,
			Comments: 43,
			Rank:     15,
		},
		{
			Title:    "post16",
			URI:      "https://www.post16.com",
			Author:   "post16Author",
			Points:   52,
			Comments: 21,
			Rank:     16,
		},
		{
			Title:    "post17",
			URI:      "https://www.post17.com",
			Author:   "post17Author",
			Points:   30,
			Comments: 14,
			Rank:     17,
		},
		{
			Title:    "post18",
			URI:      "https://www.post18.com",
			Author:   "post18Author",
			Points:   14,
			Comments: 0,
			Rank:     18,
		},
	}
)
