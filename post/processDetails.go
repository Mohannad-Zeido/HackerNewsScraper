package post

import (
	"errors"
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/Mohannad-Zeido/HackerNewsScraper/validate"
	"golang.org/x/net/html"
)

func processDetailsNode(node *html.Node) (detailsData, error) {
	points, err := getPoints(node)
	if err != nil {
		return detailsData{}, err
	}

	author, err := getAuthor(node)
	if err != nil {
		return detailsData{}, err
	}

	comments, err := getNumberOfComments(node)
	if err != nil {
		return detailsData{}, err
	}

	if !validateDetailsData(author, points, comments) {
		return detailsData{}, nil
	}

	return detailsData{
		author:   author,
		points:   points,
		comments: comments,
		valid:    true,
	}, nil
}

func getPoints(node *html.Node) (int, error) {
	pointsNode, err := getPointsNode(node)
	if err != nil {
		return 0, err
	}
	points, _ := helper.GetTagText(pointsNode)
	return helper.ExtractNumberFromString(points)
}

func getPointsNode(node *html.Node) (*html.Node, error) {
	pointsNode := helper.GetFirstChildElement(node)
	if pointsNode == nil {
		return nil, errors.New(types.ErrGettingPointsNode)
	}
	return pointsNode, nil
}

func getAuthor(node *html.Node) (string, error) {
	authorNode, err := getAuthorNode(node)
	if err != nil {
		return "", err
	}
	author, _ := helper.GetTagText(authorNode)
	return author, nil
}

func getAuthorNode(node *html.Node) (*html.Node, error) {
	detailsNode := helper.GetFirstChildElement(node)
	if detailsNode == nil {
		return nil, errors.New(types.ErrGettingAuthorDetailsNode)
	}
	authorNode := helper.GetNextSiblingElement(detailsNode)
	if authorNode == nil {
		return nil, errors.New(types.ErrGettingAuthorNode)
	}
	return authorNode, nil
}

func getNumberOfComments(node *html.Node) (int, error) {
	commentNode, err := getCommentsNode(node)
	if err != nil {
		return 0, err
	}
	commentsText, textPresent := helper.GetTagText(commentNode)
	if !textPresent {
		return -1, nil
	}
	comments, err := helper.ExtractNumberFromString(commentsText)
	if err != nil {
		return 0, err
	}
	if comments == -1 {
		comments = 0
	}
	return comments, nil
}

func getCommentsNode(node *html.Node) (*html.Node, error) {
	detailsNode := helper.GetFirstChildElement(node)
	if detailsNode == nil {
		return nil, errors.New(types.ErrGettingCommentsDetailsNode)
	}
	commentsNode := helper.GetNthSibling(detailsNode, types.CommentsNodePosition)
	if commentsNode == nil {
		return nil, errors.New(types.ErrGettingCommentsNode)
	}
	return commentsNode, nil
}

func validateDetailsData(author string, points, comments int) bool {
	if !validate.IsValidNumber(points) || !validate.IsValidNumber(comments) || !validate.IsValidText(author) {
		return false
	}
	return true
}
