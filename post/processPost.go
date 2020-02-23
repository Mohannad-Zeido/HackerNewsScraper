package post

import (
	"errors"
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"golang.org/x/net/html"
)

type generalInfoData struct {
	title string
	uri   string
	rank  int
	valid bool
}

type detailsData struct {
	author   string
	points   int
	comments int
	valid    bool
}

func getPost(currentNode *html.Node) (types.Post, bool, error) {
	var err error

	postGeneralInfo := processGeneralInfoNode(currentNode)

	if !postGeneralInfo.valid {
		return types.Post{}, false, nil
	}

	currentNode, err = getPostDetailsNode(currentNode)
	if err != nil {
		return types.Post{}, false, err
	}

	details := processDetailsNode(currentNode)
	if !details.valid {
		return types.Post{}, false, nil
	}
	return types.Post{
		Title:    postGeneralInfo.title,
		Uri:      postGeneralInfo.uri,
		Author:   details.author,
		Points:   details.points,
		Comments: details.comments,
		Rank:     postGeneralInfo.rank,
	}, true, nil
}

func getPostDetailsNode(currentNode *html.Node) (*html.Node, error) {
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
