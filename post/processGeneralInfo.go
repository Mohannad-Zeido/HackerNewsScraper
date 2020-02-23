package post

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/Mohannad-Zeido/HackerNewsScraper/validate"
	"golang.org/x/net/html"
)

func processGeneralInfoNode(node *html.Node) generalInfoData {
	rank, valid := getRank(node)
	if !valid {
		return generalInfoData{}
	}

	uri, valid := getUri(node)
	if !valid {
		return generalInfoData{}
	}

	title, valid := getTitle(node)
	if !valid {
		return generalInfoData{}
	}

	if !validateGeneralInfoData(rank, uri, title) {
		return generalInfoData{}
	}

	return generalInfoData{
		title: title,
		uri:   uri,
		rank:  rank,
		valid: true,
	}
}

func validateGeneralInfoData(rank int, uri, title string) bool {
	if !validate.IsValidNumber(rank) || !validate.IsValidURI(uri) || !validate.IsValidText(title) {
		return false
	}
	return true
}

func getTitle(node *html.Node) (string, bool) {
	titleNode, ok := getTitleNode(node)
	if !ok {
		return "", false
	}
	title, _ := helper.GetTagText(titleNode)
	return title, true
}

func getTitleNode(node *html.Node) (*html.Node, bool) {
	generalInfoNode := helper.GetFirstChildElement(node)
	if generalInfoNode == nil {
		return nil, false
	}
	nodeContainingTitleNode := helper.GetNthSibling(generalInfoNode, types.UriNodePositionInGeneralInfo)
	if nodeContainingTitleNode == nil {
		return nil, false
	}
	titleNode := helper.GetFirstChildElement(nodeContainingTitleNode)
	if titleNode == nil {
		return nil, false
	}
	return titleNode, true
}

func getUri(node *html.Node) (string, bool) {
	uriNode, ok := getUriNode(node)
	if !ok {
		return "", false
	}

	uri := helper.AttributeValue(uriNode.Attr, types.UriAttr)
	if types.InternalUriRegex.MatchString(uri) {
		uri = "https://news.ycombinator.com/" + uri
	}
	return uri, true
}

func getUriNode(node *html.Node) (*html.Node, bool) {
	generalInfoNode := helper.GetFirstChildElement(node)
	if generalInfoNode == nil {
		return nil, false
	}

	nodeContainingUriNode := helper.GetNthSibling(generalInfoNode, types.UriNodePositionInGeneralInfo)
	if nodeContainingUriNode == nil {
		return nil, false
	}
	uriNode := helper.GetFirstChildElement(nodeContainingUriNode)
	if uriNode == nil {
		return nil, false
	}
	return uriNode, true
}

func getRank(node *html.Node) (int, bool) {
	rankNode, ok := getRankNode(node)
	if !ok {
		return 0, false
	}
	rank, _ := helper.GetTagText(rankNode)
	return helper.ExtractNumberFromString(rank)
}

func getRankNode(node *html.Node) (*html.Node, bool) {
	rankNode := helper.GetChildAtDepth(node, types.RankNodeDepth)
	if rankNode == nil {
		return nil, false
	}
	return rankNode, true
}
