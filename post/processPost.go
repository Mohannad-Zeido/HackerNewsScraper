package post

import (
	"errors"
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"golang.org/x/net/html"
)

//generalInfoData refers to the information in the first row of a post
type generalInfoData struct {
	title string
	uri   string
	rank  int
	valid bool
}

//detailsData refers to the information in the second row of a post
type detailsData struct {
	author   string
	points   int
	comments int
	valid    bool
}

//processPost will parse the data in each row of the post and if all the information is valid will return true and the
//processed post. This function will return an error only if the details node of a post can not be retrieved
func processPost(currentNode *html.Node) (types.Post, bool, error) {
	var err error

	postGeneralInfo := processGeneralInfoNode(currentNode)

	if !postGeneralInfo.valid {
		return types.Post{}, false, nil
	}

	currentNode, err = getDetailsNode(currentNode)
	if err != nil {
		return types.Post{}, false, err
	}

	details := processDetailsNode(currentNode)
	if !details.valid {
		return types.Post{}, false, nil
	}
	return types.Post{
		Title:    postGeneralInfo.title,
		URI:      postGeneralInfo.uri,
		Author:   details.author,
		Points:   details.points,
		Comments: details.comments,
		Rank:     postGeneralInfo.rank,
	}, true, nil
}

//getDetailsNode will traverse the post nodes and return the node that contains the second row details of a post
//An error will be returned if the post nodes are not in the expected structure
func getDetailsNode(currentNode *html.Node) (*html.Node, error) {
	currentNode = helper.GetNextSiblingElement(currentNode)
	if currentNode == nil {
		return nil, errors.New(types.ErrGettingPostDetailsNode)
	}

	currentNode = helper.GetFirstChildElement(currentNode)
	if currentNode == nil {
		return nil, errors.New(types.ErrGettingPostDetailsNode)
	}

	currentNode = helper.GetNextSiblingElement(currentNode)
	if currentNode == nil || !helper.ContainsAttributeValue(currentNode.Attr, types.DetailsTag) {
		return nil, errors.New(types.ErrGettingPostDetailsNode)
	}
	return currentNode, nil
}
