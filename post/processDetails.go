package post

import (
	"errors"
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/Mohannad-Zeido/HackerNewsScraper/validate"
	"golang.org/x/net/html"
	"log"
)

func processDetailsNode(node *html.Node) detailsData {
	points, valid := getPoints(node)
	if !valid {
		return detailsData{}
	}

	author, valid := getAuthor(node)
	if !valid {
		return detailsData{}
	}

	comments, valid := getNumberOfComments(node)
	if !valid {
		return detailsData{}
	}

	if !validateDetailsData(author, points, comments) {
		return detailsData{}
	}

	return detailsData{
		author:   author,
		points:   points,
		comments: comments,
		valid:    true,
	}
}

func getPoints(node *html.Node) (int, bool) {
	pointsNode, err := getPointsNode(node)
	if err != nil {
		log.Println(err)
		return 0, false
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

func getAuthor(node *html.Node) (string, bool) {
	authorNode, err := getAuthorNode(node)
	if err != nil {
		log.Println(err)
		return "", false
	}
	return helper.GetTagText(authorNode)
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

func getNumberOfComments(node *html.Node) (int, bool) {
	commentNode, err := getCommentsNode(node)
	if err != nil {
		log.Println(err)
		return 0, false
	}
	commentsText, textPresent := helper.GetTagText(commentNode)
	if !textPresent {
		return -1, false
	}
	comments, validInt := helper.ExtractNumberFromString(commentsText)
	if !validInt {
		return 0, validInt
	}
	if comments == -1 {
		comments = 0
	}
	return comments, true
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
