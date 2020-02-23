package post

import (
	"errors"
	"fmt"
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/parse"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"golang.org/x/net/html"
)

func GetPosts(numPosts int) ([]types.Post, error) {
	currentPage := 0
	postsLeftToGet := numPosts
	var posts []types.Post
	for {
		currentPage += 1
		pageNode, err := parse.GetData(currentPage)
		if err != nil {
			return nil, err
		}
		postsFromPage, err := getPostsFromPage(pageNode, postsLeftToGet)
		if err != nil {
			return nil, err
		}
		posts = append(posts, postsFromPage...)
		postsLeftToGet = postsLeftToGet - len(postsFromPage)
		if postsLeftToGet == 0 {
			return posts, nil
		}
	}
}

func getPostsFromPage(node *html.Node, numPosts int) ([]types.Post, error) {
	var result []types.Post

	currentNode, err := findFirstPostNode(node)
	if err != nil || !helper.ContainsAttributeValue(currentNode.Attr, types.GeneralInfoTag) {
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
	postNode := helper.GetNthSibling(node, types.NumberOfNodesPerPost)
	if postNode == nil {
		return nil, errors.New(types.ErrGettingNextPost)
	}
	if !helper.ContainsAttributeValue(postNode.Attr, types.GeneralInfoTag) {
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
	tableNode := helper.TagFinder(node, types.TableTag, types.PostsTableAttrVal)
	if tableNode == nil {
		return nil, fmt.Errorf(types.ErrParsingHTML)
	}
	return tableNode, nil
}

func getFirstRecordInTable(tableNode *html.Node) (*html.Node, error) {
	tBodyNode := helper.GetFirstChildElement(tableNode)
	if tBodyNode.Data != types.TbodyTag {
		return nil, fmt.Errorf(types.ErrParsingHTML)
	}
	firstRecordNode := helper.GetFirstChildElement(tBodyNode)
	if firstRecordNode == nil {
		return nil, fmt.Errorf(types.ErrParsingHTML)
	}
	return firstRecordNode, nil
}
