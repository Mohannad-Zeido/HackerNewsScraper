package post

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/helper"
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"github.com/Mohannad-Zeido/HackerNewsScraper/validate"
	"golang.org/x/net/html"
	"regexp"
)

//regex that will be used to match the internal page link
var internalURIRegex, _ = regexp.Compile(types.InternalURI)

//processGeneralInfoNode will extract and validate all the information in the first row of the post
func processGeneralInfoNode(node *html.Node) generalInfoData {
	rank, ok := getRank(node)
	if !ok {
		return generalInfoData{}
	}

	uri, ok := getURI(node)
	if !ok {
		return generalInfoData{}
	}

	title, ok := getTitle(node)
	if !ok {
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

//getTitle will return the title for the current post
func getTitle(node *html.Node) (string, bool) {
	titleNode, ok := getTitleNode(node)
	if !ok {
		return "", false
	}
	title, _ := helper.GetTagText(titleNode)
	return title, true
}

//getTitleNode is a wrapper to the getUriTitleNode
func getTitleNode(node *html.Node) (*html.Node, bool) {
	return getUriTitleNode(node)
}

//getURI will return the URI of the current post
func getURI(node *html.Node) (string, bool) {
	uriNode, ok := getURINode(node)
	if !ok {
		return "", false
	}

	uri := helper.AttributeValue(uriNode.Attr, types.URIAttr)
	//if the uri is pointing to an internal page build the whole uri
	if internalURIRegex.MatchString(uri) {
		uri = "https://news.ycombinator.com/" + uri
	}

	return uri, true
}

//getURINode is a wrapper to the getUriTitleNode
func getURINode(node *html.Node) (*html.Node, bool) {
	return getUriTitleNode(node)
}

//getUriTitleNode will return the node that contains both the URI and Title. This function assumes the parameter is
//pointing to the posts' first row node.
func getUriTitleNode(node *html.Node) (*html.Node, bool) {
	generalInfoNode := helper.GetFirstChildElement(node)
	if generalInfoNode == nil {
		return nil, false
	}

	uriTitleParentNode := helper.GetNthSibling(generalInfoNode, types.URINodePositionInGeneralInfo)
	if uriTitleParentNode == nil {
		return nil, false
	}
	uriTitleNode := helper.GetFirstChildElement(uriTitleParentNode)
	if uriTitleNode == nil {
		return nil, false
	}
	return uriTitleNode, true

}

//getRank will return the Rank of the current post
func getRank(node *html.Node) (int, bool) {
	rankNode, ok := getRankNode(node)
	if !ok {
		return 0, false
	}
	rank, _ := helper.GetTagText(rankNode)
	return helper.ExtractNumberFromString(rank)
}

//getRankNode will return the node that contains the rank. This function assumes the parameter is
//pointing to the posts' first row node.
func getRankNode(node *html.Node) (*html.Node, bool) {
	rankNode := helper.GetChildAtDepth(node, types.RankNodeDepth)
	if rankNode == nil {
		return nil, false
	}
	return rankNode, true
}
