package post

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/Mohannad-Zeido/HackerNewsScraper/validate"
	"golang.org/x/net/html"
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
	pointsNode, ok := getPointsNode(node)
	if !ok {
		return 0, false
	}
	points, _ := helper.GetTagText(pointsNode)
	return helper.ExtractNumberFromString(points)
}

func getPointsNode(node *html.Node) (*html.Node, bool) {
	pointsNode := helper.GetFirstChildElement(node)
	if pointsNode == nil {
		return nil, false
	}
	return pointsNode, true
}

func getAuthor(node *html.Node) (string, bool) {
	authorNode, ok := getAuthorNode(node)
	if !ok {
		return "", false
	}
	return helper.GetTagText(authorNode)
}

func getAuthorNode(node *html.Node) (*html.Node, bool) {
	detailsNode := helper.GetFirstChildElement(node)
	if detailsNode == nil {
		return nil, false
	}
	authorNode := helper.GetNextSiblingElement(detailsNode)
	if authorNode == nil {
		return nil, false
	}
	return authorNode, true
}

func getNumberOfComments(node *html.Node) (int, bool) {
	commentNode, ok := getCommentsNode(node)
	if !ok {
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

func getCommentsNode(node *html.Node) (*html.Node, bool) {
	detailsNode := helper.GetFirstChildElement(node)
	if detailsNode == nil {
		return nil, false
	}
	commentsNode := helper.GetNthSibling(detailsNode, types.CommentsNodePosition)
	if commentsNode == nil {
		return nil, false
	}
	return commentsNode, true
}

func validateDetailsData(author string, points, comments int) bool {
	if !validate.IsValidNumber(points) || !validate.IsValidNumber(comments) || !validate.IsValidText(author) {
		return false
	}
	return true
}
