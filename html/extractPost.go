package html

import (
	"fmt"
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
	errParsingHTML    = "error scraping the web page"
)

var (
	nonNumbersRegex, _  = regexp.Compile(nonNumbers)
	internalUriRegex, _ = regexp.Compile(internalUri)
)

// todo i need to return error rather than nil when it is actually a parsing/traversal error
func GetPosts(numPosts int) []types.Post {
	//todo return errors
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
		postsFromPage, err := getPostNodes(pageNode, postsLeftToGet)
		if err != nil {
			panic(err)
		}
		posts = append(posts, postsFromPage...)
		postsLeftToGet = postsLeftToGet - len(postsFromPage)
		if postsLeftToGet == 0 {
			return posts
		}
	}

}

func getPostNodes(node *html.Node, numPosts int) ([]types.Post, error) {
	var result []types.Post

	child, err := findFirstPostNode(node)
	if err != nil {
		return nil, err
	}

	for {
		var post types.Post

		post, postValid, err := extractPost(child)
		if err != nil {
			return nil, err
		}

		if postValid {
			result = append(result, post)
		}

		child, err = getNextPost(child)
		if err != nil {
			return nil, err
		}

		if len(result) == numPosts {
			return result, nil
		}
	}
}

func getNextPost(currentNode *html.Node) (*html.Node, error) {

	currentNode = getNextSiblingElementNode(currentNode)
	if currentNode == nil {
		return nil, fmt.Errorf(errParsingHTML)
	}
	currentNode = getNextSiblingElementNode(currentNode)
	if currentNode == nil {
		return nil, fmt.Errorf(errParsingHTML)
	}
	currentNode = getNextSiblingElementNode(currentNode)
	if currentNode == nil || containsAttributeValue(currentNode.Attr, endOfPostsAttrVal) {
		return nil, fmt.Errorf(errParsingHTML)
	}
	return currentNode, nil
}

func extractPost(currentNode *html.Node) (types.Post, bool, error) {
	var post types.Post
	var postValid bool
	var err error

	if currentNode == nil || !containsAttributeValue(currentNode.Attr, generalInfoTag) {
		return types.Post{}, false, fmt.Errorf(errParsingHTML)
	}
	post.Title, post.Uri, post.Rank, postValid = parseGeneralInfoNode(currentNode)
	if !postValid {
		return types.Post{}, false, nil
	}
	currentNode, err = getPostDetailsNode(currentNode)
	if err != nil {
		return types.Post{}, false, fmt.Errorf(errParsingHTML)
	}

	post.Author, post.Comments, post.Points, postValid = parseDetailsNode(currentNode)
	if !postValid {
		return types.Post{}, false, nil
	}
	return post, true, nil
}

func getPostDetailsNode(currentNode *html.Node) (*html.Node, error) {
	currentNode = getNextSiblingElementNode(currentNode)
	if currentNode == nil {
		return nil, fmt.Errorf(errParsingHTML)
	}
	currentNode = getFirstChildElementNode(currentNode)
	if currentNode == nil {
		return nil, fmt.Errorf(errParsingHTML)
	}
	currentNode = getNextSiblingElementNode(currentNode)
	if currentNode == nil || !containsAttributeValue(currentNode.Attr, detailsTag) {
		return nil, fmt.Errorf(errParsingHTML)
	}
	return currentNode, nil
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

func findFirstPostNode(node *html.Node) (*html.Node, error) {
	tableNode, err := findTableOfPosts(node)
	if err != nil {
		return nil, err
	}
	return getFirstRecordInTable(tableNode)
}

func findTableOfPosts(node *html.Node) (*html.Node, error) {
	tableNode := tagFinder(node, tableTag, postsTableAttrVal)
	if tableNode == nil {
		return nil, fmt.Errorf(errParsingHTML)
	}
	return tableNode, nil
}

func getFirstRecordInTable(tableNode *html.Node) (*html.Node, error) {
	tBodyNode := getFirstChildElementNode(tableNode)
	if tBodyNode.Data != tbodyTag {
		return nil, fmt.Errorf(errParsingHTML)
	}
	firstRecordNode := getFirstChildElementNode(tBodyNode)
	if firstRecordNode == nil {
		return nil, fmt.Errorf(errParsingHTML)
	}
	return firstRecordNode, nil
}
