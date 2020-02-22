package html

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/parse"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/Mohannad-Zeido/HackerNewsScraper/validate"
	"golang.org/x/net/html"
	"regexp"
	"strconv"
)

const (
	tableTag          = "table"
	generalInfoTag    = "athing"
	postsTableAttrVal = "itemlist"
	tbodyTag          = "tbody"
	hrefAttr          = "href"
	detailsTag        = "subtext"
	endOfPostsAttrVal = "morespace"
	nonNumbers        = "\\D"
	internalUri       = "^item\\?id=[0-9a-zA-Z]*"
)

var (
	nonNumbersRegex, _  = regexp.Compile(nonNumbers)
	internalUriRegex, _ = regexp.Compile(internalUri)
)

// todo i need to return error rather than nil when it is actually a parsing/traversal error
func GetPosts(numPosts int) []types.Post {
	//data, err := parse.GetData("data.html")
	currentPage := 0
	postsLeftToGet := numPosts
	var posts []types.Post
	for {
		currentPage += 1
		pageNode, err := parse.GetData(currentPage)
		if err != nil {
			panic(err)
		}
		postsFromPage := getPostNodes(pageNode, postsLeftToGet)
		posts = append(posts, postsFromPage...)
		postsLeftToGet = postsLeftToGet - len(postsFromPage)
		if postsLeftToGet == 0 {
			return posts
		}
	}

}

func getPostNodes(node *html.Node, numPosts int) []types.Post {
	var result []types.Post
	child := findPosts(node)
	if child == nil {
		return nil
	}
	child = getFirstChildElementNode(child)
	if child == nil {
		return nil
	}

	for {
		var post types.Post

		post, postValid := extractPost(child)
		if postValid {
			result = append(result, post)
		}

		child = getNextPost(child)
		if child == nil || len(result) == numPosts {
			return result
		}
	}
}

func getNextPost(currentNode *html.Node) *html.Node {

	currentNode = getNextSiblingElementNode(currentNode)
	if currentNode == nil {
		return nil
	}
	currentNode = getNextSiblingElementNode(currentNode)
	if currentNode == nil {
		return nil
	}
	currentNode = getNextSiblingElementNode(currentNode)
	if currentNode == nil || containsAttributeValue(currentNode.Attr, endOfPostsAttrVal) {
		return nil
	}
	return currentNode
}

func extractPost(currentNode *html.Node) (types.Post, bool) {
	var post types.Post
	var postValid bool

	if currentNode == nil || !containsAttributeValue(currentNode.Attr, generalInfoTag) {
		//todo return error
		return types.Post{}, false
	}
	post.Title, post.Uri, post.Rank, postValid = parseGeneralInfoNode(currentNode)
	if !postValid {
		return types.Post{}, false
	}
	currentNode = getGeneralInfoNode(currentNode)
	if currentNode == nil || !containsAttributeValue(currentNode.Attr, detailsTag) {
		//todo return error
		return types.Post{}, false
	}
	post.Author, post.Comments, post.Points, postValid = parseDetailsNode(currentNode)

	if !postValid {
		return types.Post{}, false
	}
	return post, true
}

func getGeneralInfoNode(currentNode *html.Node) *html.Node {
	currentNode = getNextSiblingElementNode(currentNode)
	if currentNode == nil {
		//todo return error
		return nil
	}
	currentNode = getFirstChildElementNode(currentNode)
	if currentNode == nil {
		//todo return error
		return nil
	}
	return getNextSiblingElementNode(currentNode)
}

func parseGeneralInfoNode(node *html.Node) (string, string, int, bool) {
	//todo do some null child error checking
	rankTD := getFirstChildElementNode(node)
	spanRank := getFirstChildElementNode(rankTD)
	//todo if firstChild is nil invalid post
	rank, err := extractNumber(tagText(spanRank))
	if err != nil || !validate.IsValidNumber(rank) {
		return "", "", 0, false
	}
	href := getFirstChildElementNode(getNextSiblingElementNode(getNextSiblingElementNode(rankTD)))
	uri := attributeValue(href.Attr, hrefAttr)
	if internalUriRegex.MatchString(uri) {
		uri = "https://news.ycombinator.com/" + uri
	}
	if !validate.IsValidUri(uri) {
		return "", "", 0, false
	}
	title := tagText(href)
	if !validate.IsValidText(title) {
		return "", "", 0, false
	}
	return title, uri, rank, true
}

func extractNumber(s string) (int, error) {
	n, err := strconv.Atoi(nonNumbersRegex.ReplaceAllString(s, ""))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func parseDetailsNode(node *html.Node) (string, int, int, bool) {
	scoreSpan := getFirstChildElementNode(node)
	//score := tagText(scoreSpan)
	userA := getNextSiblingElementNode(scoreSpan)
	user := tagText(userA)
	//commentsA := getNextSiblingElementNode(getNextSiblingElementNode(getNextSiblingElementNode(getNextSiblingElementNode(userA))))
	//comments := tagText(commentsA)
	return user, 1, 4, true
}

func findPosts(node *html.Node) *html.Node {
	table := tagFinder(node, tableTag, postsTableAttrVal)
	if table == nil {
		return nil
	}

	tb := getFirstChildElementNode(table)
	if tb.Data != tbodyTag {
		return nil
	}
	return tb
}
