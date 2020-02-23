package post

import (
	"errors"
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/Mohannad-Zeido/HackerNewsScraper/validate"
	"golang.org/x/net/html"
	"log"
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
	if !validate.IsValidNumber(rank) || !validate.IsValidUri(uri) || !validate.IsValidText(title) {
		return false
	}
	return true
}

func getTitle(node *html.Node) (string, bool) {
	titleNode, err := getTitleNode(node)
	if err != nil {
		log.Println(err)
		return "", false
	}
	title, _ := helper.GetTagText(titleNode)
	return title, true
}

func getTitleNode(node *html.Node) (*html.Node, error) {
	generalInfoNode := helper.GetFirstChildElement(node)
	if generalInfoNode == nil {
		return nil, errors.New(types.ErrGettingTitleGeneralInfoNode)
	}
	nodeContainingTitleNode := helper.GetNthSibling(generalInfoNode, types.UriNodePositionInGeneralInfo)
	if nodeContainingTitleNode == nil {
		return nil, errors.New(types.ErrGettingTitleParentNode)
	}
	titleNode := helper.GetFirstChildElement(nodeContainingTitleNode)
	if titleNode == nil {
		return nil, errors.New(types.ErrGettingTitleNode)
	}
	return titleNode, nil
}

func getUri(node *html.Node) (string, bool) {
	uriNode, err := getUriNode(node)
	if err != nil {
		log.Println(err)
		return "", false
	}

	uri := helper.AttributeValue(uriNode.Attr, types.UriAttr)
	if types.InternalUriRegex.MatchString(uri) {
		uri = "https://news.ycombinator.com/" + uri
	}
	return uri, true
}

func getUriNode(node *html.Node) (*html.Node, error) {
	generalInfoNode := helper.GetFirstChildElement(node)
	if generalInfoNode == nil {
		return nil, errors.New(types.ErrGettingUriGeneralInfoNode)
	}

	nodeContainingUriNode := helper.GetNthSibling(generalInfoNode, types.UriNodePositionInGeneralInfo)
	if nodeContainingUriNode == nil {
		return nil, errors.New(types.ErrGettingUriParentNode)
	}
	uriNode := helper.GetFirstChildElement(nodeContainingUriNode)
	if uriNode == nil {
		return nil, errors.New(types.ErrGettingUriNode)
	}
	return uriNode, nil
}

func getRank(node *html.Node) (int, bool) {
	rankNode, err := getRankNode(node)
	if err != nil {
		log.Println(err)
		return 0, false
	}
	rank, _ := helper.GetTagText(rankNode)
	return helper.ExtractNumberFromString(rank)
}

func getRankNode(node *html.Node) (*html.Node, error) {
	rankNode := helper.GetChildAtDepth(node, types.RankNodeDepth)
	if rankNode == nil {
		return nil, errors.New(types.ErrGettingRankNode)
	}
	return rankNode, nil
}
