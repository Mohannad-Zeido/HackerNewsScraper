package post

import (
	"errors"
	"fmt"
	"github.com/Mohannad-Zeido/HackerNewsScraper/dateRetriever"
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"golang.org/x/net/html"
)

//GetPosts will return a list of posts that have been retrieved from the data source
//while the number of posts to get is not 0 this function will loop through the pages available in the data source until
//enough posts have been retrieved. An error in the retrieval of data or HTML traversal will break the loop
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

//getPostsFromPage will return the valid posts on a certain page.
//this function will either either get the number of valid posts required as indicated with the numPosts parameter of
//all the valid posts on the page if the number of posts required has not yet been reached.
func getPostsFromPage(node *html.Node, numPosts int) ([]types.Post, error) {
	var result []types.Post

	currentNode, err := findFirstPostNode(node)
	if err != nil {
		return nil, err
	}

	for {
		var post types.Post

		post, postValid, err := processPost(currentNode)
		if err != nil {
			return nil, err
		}

		//only valid posts will be added to the results
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

//getNextPost will traverse the posts table returning a pointer to the start of the next post
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

//firstRecordNode will travers the page from the currentLocation of the node (usually the beginning of the page and
//return a pointer to first post in the table
func findFirstPostNode(node *html.Node) (*html.Node, error) {
	tableNode, err := findTableOfPosts(node)
	if err != nil {
		return nil, err
	}
	return getFirstRecordInTable(tableNode)
}

//findTableOfPosts will return the table node that contains all the posts on the page
func findTableOfPosts(node *html.Node) (*html.Node, error) {
	tableNode := helper.TagFinder(node, types.TableTag, types.PostsTableAttrVal)
	if tableNode == nil {
		return nil, fmt.Errorf(types.ErrGettingPostsTableNode)
	}
	return tableNode, nil
}

//getFirstRecordInTable will assume the node passed in is a table node and will return the first row of that table if
//that row is a post as indicated by the class attribute
func getFirstRecordInTable(tableNode *html.Node) (*html.Node, error) {
	//skip tbody Tag
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
