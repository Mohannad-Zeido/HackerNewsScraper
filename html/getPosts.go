package html

import (
	"errors"
	"fmt"
	"github.com/Mohannad-Zeido/HackerNewsScraper/parse"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"golang.org/x/net/html"
	"regexp"
)

const (
	tableTag                       = "table"
	generalInfoTag                 = "athing"
	postsTableAttrVal              = "itemlist"
	tbodyTag                       = "tbody"
	uriAttr                        = "href"
	detailsTag                     = "subtext"
	endOfPostsAttrVal              = "morespace"
	nonNumbers                     = "\\D"
	internalUri                    = "^item\\?id=[0-9a-zA-Z]*"
	numberOfNodesPerPost           = 3
	uriNodePositionInGeneralInfo   = 2
	tableNodePositionInGeneralInfo = 3
	rankNodeDepth                  = 2
	commentsNodePosition           = 5
)

var (
	nonNumbersRegex, _  = regexp.Compile(nonNumbers)
	internalUriRegex, _ = regexp.Compile(internalUri)
)

// todo i need to return error rather than nil when it is actually a parsing/traversal error
func GetPosts(numPosts int) []types.Post {
	//todo return errors instead of panic
	currentPage := 0
	postsLeftToGet := numPosts
	var posts []types.Post
	for {
		currentPage += 1
		pageNode, err := parse.GetData(currentPage)
		if err != nil {
			panic(err)
		}
		postsFromPage, err := getPostsFromPage(pageNode, postsLeftToGet)
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

func getPostsFromPage(node *html.Node, numPosts int) ([]types.Post, error) {
	var result []types.Post

	currentNode, err := findFirstPostNode(node)
	if err != nil || !containsAttributeValue(currentNode.Attr, generalInfoTag) {
		return nil, err
	}

	for {
		var post types.Post

		post, postValid, err := getPost(currentNode)
		if err != nil {
			return nil, err
		}

		if postValid {
			result = append(result, post)
		}
		currentNode, err = getNextPost(currentNode)
		if err != nil {
			return nil, err
		}

		if currentNode == nil || len(result) >= numPosts {
			return result, nil
		}
	}
}

func getNextPost(node *html.Node) (*html.Node, error) {
	postNode := getNthSibling(node, numberOfNodesPerPost)
	if postNode == nil {
		return nil, errors.New(types.ErrGettingNextPost)
	}
	if !containsAttributeValue(postNode.Attr, generalInfoTag) {
		return nil, nil
	}
	return postNode, nil
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
		return nil, fmt.Errorf(types.ErrParsingHTML)
	}
	return tableNode, nil
}

func getFirstRecordInTable(tableNode *html.Node) (*html.Node, error) {
	tBodyNode := getFirstChildElementNode(tableNode)
	if tBodyNode.Data != tbodyTag {
		return nil, fmt.Errorf(types.ErrParsingHTML)
	}
	firstRecordNode := getFirstChildElementNode(tBodyNode)
	if firstRecordNode == nil {
		return nil, fmt.Errorf(types.ErrParsingHTML)
	}
	return firstRecordNode, nil
}
