package post

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/Mohannad-Zeido/HackerNewsScraper/validate"
	"golang.org/x/net/html"
)

//processDetailsNode will extract and validate all the information in the second row of the post
func processDetailsNode(node *html.Node) detailsData {
	points, ok := getPoints(node)
	if !ok {
		return detailsData{}
	}

	author, ok := getAuthor(node)
	if !ok {
		return detailsData{}
	}

	comments, ok := getNumberOfComments(node)
	if !ok {
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

func validateDetailsData(author string, points, comments int) bool {
	if !validate.IsValidNumber(points) || !validate.IsValidNumber(comments) || !validate.IsValidText(author) {
		return false
	}
	return true
}

//getPoints will return the points for the current post
func getPoints(node *html.Node) (int, bool) {
	pointsNode, ok := getPointsNode(node)
	if !ok {
		return 0, false
	}
	points, _ := helper.GetTagText(pointsNode)
	return helper.ExtractNumberFromString(points)
}

//getPointsNode will return the node that contains the points. This function assumes the parameter is
//pointing to the posts' second row node.
func getPointsNode(node *html.Node) (*html.Node, bool) {
	pointsNode := helper.GetFirstChildElement(node)
	if pointsNode == nil {
		return nil, false
	}
	return pointsNode, true
}

//getAuthor will return the author of the current post
func getAuthor(node *html.Node) (string, bool) {
	authorNode, ok := getAuthorNode(node)
	if !ok {
		return "", false
	}
	return helper.GetTagText(authorNode)
}

//getAuthorNode will return the node that contains the author. This function assumes the parameter is
//pointing to the posts' second row node.
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

//getNumberOfComments will return the number of comments for the current post
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

//getCommentsNode will return the node that contains the comments. This function assumes the parameter is
//pointing to the posts' second row node.
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
