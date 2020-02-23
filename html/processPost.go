package html

import (
	"errors"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/Mohannad-Zeido/HackerNewsScraper/validate"
	"golang.org/x/net/html"
	"strconv"
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

	postGeneralInfo, err := parseGeneralInfoNode(currentNode)
	if err != nil {
		return types.Post{}, false, err
	}
	if !postGeneralInfo.valid {
		return types.Post{}, false, nil
	}

	currentNode, err = getPostDetailsNode(currentNode)
	if err != nil {
		return types.Post{}, false, err
	}

	details, err := parseDetailsNode(currentNode)
	if err != nil {
		return types.Post{}, false, err
	}

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
	currentNode = getNextSiblingElementNode(currentNode)
	if currentNode == nil {
		return nil, errors.New(types.ErrGettingPostDetailsNode)
	}
	currentNode = getFirstChildElementNode(currentNode)
	if currentNode == nil {
		return nil, errors.New(types.ErrGettingPostDetailsNode)
	}
	currentNode = getNextSiblingElementNode(currentNode)
	if currentNode == nil || !containsAttributeValue(currentNode.Attr, detailsTag) {
		return nil, errors.New(types.ErrGettingPostDetailsNode)
	}
	return currentNode, nil
}

func parseGeneralInfoNode(node *html.Node) (generalInfoData, error) {
	rank, err := getRank(node)
	if err != nil {
		return generalInfoData{}, err
	}

	uri, err := getUri(node)
	if err != nil {
		return generalInfoData{}, err
	}

	title, err := getTitle(node)
	if err != nil {
		return generalInfoData{}, err
	}

	if !validateGeneralInfoData(rank, uri, title) {
		return generalInfoData{}, nil
	}

	return generalInfoData{
		title: title,
		uri:   uri,
		rank:  rank,
		valid: true,
	}, nil
}

func parseDetailsNode(node *html.Node) (detailsData, error) {
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

func validateGeneralInfoData(rank int, uri, title string) bool {
	if !validate.IsValidNumber(rank) || !validate.IsValidUri(uri) || !validate.IsValidText(title) {
		return false
	}
	return true
}

func getTitle(node *html.Node) (string, error) {
	titleNode, err := getTitleNode(node)
	if err != nil {
		return "", err
	}
	return tagText(titleNode), nil
}

func getTitleNode(node *html.Node) (*html.Node, error) {
	generalInfoNode := getFirstChildElementNode(node)
	if generalInfoNode == nil {
		return nil, errors.New(types.ErrGettingTitleGeneralInfoNode)
	}
	nodeContainingTitleNode := getNthSibling(generalInfoNode, uriNodePositionInGeneralInfo)
	if nodeContainingTitleNode == nil {
		return nil, errors.New(types.ErrGettingTitleParentNode)
	}
	titleNode := getFirstChildElementNode(nodeContainingTitleNode)
	if titleNode == nil {
		return nil, errors.New(types.ErrGettingTitleNode)
	}
	return titleNode, nil
}

func getUri(node *html.Node) (string, error) {
	uriNode, err := getUriNode(node)
	if err != nil {
		return "", err
	}

	uri := attributeValue(uriNode.Attr, uriAttr)
	if internalUriRegex.MatchString(uri) {
		uri = "https://news.ycombinator.com/" + uri
	}
	return uri, err
}

func getUriNode(node *html.Node) (*html.Node, error) {
	generalInfoNode := getFirstChildElementNode(node)
	if generalInfoNode == nil {
		return nil, errors.New(types.ErrGettingUriGeneralInfoNode)
	}

	nodeContainingUriNode := getNthSibling(generalInfoNode, uriNodePositionInGeneralInfo)
	if nodeContainingUriNode == nil {
		return nil, errors.New(types.ErrGettingUriParentNode)
	}
	uriNode := getFirstChildElementNode(nodeContainingUriNode)
	if uriNode == nil {
		return nil, errors.New(types.ErrGettingUriNode)
	}
	return uriNode, nil
}

func getRank(node *html.Node) (int, error) {
	rankNode, err := getRankNode(node)
	if err != nil {
		return 0, err
	}
	return extractNumberFromString(tagText(rankNode))
}

func getRankNode(node *html.Node) (*html.Node, error) {
	rankNode := getChildAtDepth(node, rankNodeDepth)
	if rankNode == nil {
		return nil, errors.New(types.ErrGettingRankNode)
	}
	return rankNode, nil
}

func extractNumberFromString(s string) (int, error) {
	number := nonNumbersRegex.ReplaceAllString(s, "")
	if number == "" {
		number = "-1"
	}
	n, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func getPoints(node *html.Node) (int, error) {
	pointsNode, err := getPointsNode(node)
	if pointsNode != nil {
		return 0, err
	}
	return extractNumberFromString(tagText(pointsNode))
}

func getPointsNode(node *html.Node) (*html.Node, error) {
	pointsNode := getFirstChildElementNode(node)
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
	return tagText(authorNode), nil
}

func getAuthorNode(node *html.Node) (*html.Node, error) {
	detailsNode := getFirstChildElementNode(node)
	if detailsNode == nil {
		return nil, errors.New(types.ErrGettingAuthorDetailsNode)
	}
	authorNode := getNextSiblingElementNode(detailsNode)
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
	comments, err := extractNumberFromString(tagText(commentNode))
	if err != nil {
		return 0, err
	}
	if comments == -1 {
		comments = 0
	}
	return comments, nil
}

func getCommentsNode(node *html.Node) (*html.Node, error) {
	detailsNode := getFirstChildElementNode(node)
	if detailsNode == nil {
		return nil, errors.New(types.ErrGettingCommentsDetailsNode)
	}
	commentsNode := getNthSibling(detailsNode, commentsNodePosition)
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
