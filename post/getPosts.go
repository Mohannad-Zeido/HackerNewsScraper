package post

import (
	"errors"
	"fmt"
	"github.com/Mohannad-Zeido/HackerNewsScraper/dateRetriever"
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"golang.org/x/net/html"
)

func GetPosts(numPosts int, state types.RunState) ([]types.Post, error) {
	if numPosts == 0 {
		return []types.Post{}, nil
	}
	var currentPage int
	postsLeftToGet := numPosts
	var posts []types.Post
	for {
		currentPage += 1
		pageNode, err := dateRetriever.GetData(currentPage, state)
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
	if err != nil {
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
	if helper.ContainsAttributeValue(postNode.Attr, types.EndOfPostsAttrVal) {
		return nil, nil
	}
	if !helper.ContainsAttributeValue(postNode.Attr, types.GeneralInfoTag) {
		return nil, errors.New(types.ErrGettingNextPost)
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
		return nil, fmt.Errorf(types.ErrGettingPostsTableNode)
	}
	return tableNode, nil
}

func getFirstRecordInTable(tableNode *html.Node) (*html.Node, error) {
	tBodyNode := helper.GetFirstChildElement(tableNode)
	if tBodyNode == nil || tBodyNode.Data != types.TbodyTag {
		return nil, fmt.Errorf(types.ErrGettingPostsTbodyNode)
	}

	firstRecordNode := helper.GetFirstChildElement(tBodyNode)
	if firstRecordNode == nil || !helper.ContainsAttributeValue(firstRecordNode.Attr, types.GeneralInfoTag) {
		return nil, fmt.Errorf(types.ErrNoPostsOnPage)
	}
	return firstRecordNode, nil
}
