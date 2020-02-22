package html

import (
	"fmt"
	"github.com/Mohannad-Zeido/HackerNewsScraper/parse"
	"golang.org/x/net/html"
)

type postNode struct {
	generalInfo *html.Node
	details     *html.Node
}

const (
	tableTag          = "table"
	generalInfoTag    = "athing"
	postsTableAttrVal = "itemlist"
	tbodyTag          = "tbody"
	hrefAttr          = "href"
	detailsTag        = "subtext"
)

func GetPosts(numPosts int) {
	data, err := parse.GetData("data.html")
	if err != nil {
		panic(err)
	}
	getPostNodes(data)
}

func getPostNodes(node *html.Node) []postNode {
	var result []postNode
	child := findPosts(node)
	child = getFirstChildElementNode(child)

	for {
		var post postNode

		if child == nil || !containsAttributeValue(child.Attr, generalInfoTag) {
			return result
		}
		post.generalInfo = child
		_, _, _, _ = parseGeneralInfo(child)
		//todo validate the info from general info
		child = getNextSiblingElementNode(child)
		if child == nil {
			//todo return error
			return result
		}
		child = getFirstChildElementNode(child)
		if child == nil {
			//todo return error
			return result
		}
		child = getNextSiblingElementNode(child)
		if child == nil || !containsAttributeValue(child.Attr, detailsTag) {
			//todo return error
			return result
		}
		post.details = child
		_, _, _, _ = parseDetails(child)
		result = append(result, post)
		child = getNextSiblingElementNode(child)
		if child == nil {
			//todo return error
			return result
		}
		child = getNextSiblingElementNode(child)
	}
}

func parseGeneralInfo(node *html.Node) (string, string, int, error) {
	//todo do some null child error checking
	rankTD := getFirstChildElementNode(node)
	spanRank := getFirstChildElementNode(rankTD)
	//todo if firstChild is nil invalid post
	rank := spanRank.FirstChild.Data
	titleTD := getNextSiblingElementNode(getNextSiblingElementNode(rankTD))
	href := titleTD.FirstChild
	uri := attributeValue(href.Attr, hrefAttr)
	hrefText := href.FirstChild.Data
	fmt.Println(hrefText)
	fmt.Println(uri)
	fmt.Println(rank)
	return "", "", 0, nil
}

func parseDetails(node *html.Node) (string, int, int, error) {
	scoreSpan := getFirstChildElementNode(node)
	score := scoreSpan.FirstChild.Data
	userA := getNextSiblingElementNode(scoreSpan)
	user := userA.FirstChild.Data
	commentsA := getNextSiblingElementNode(getNextSiblingElementNode(getNextSiblingElementNode(getNextSiblingElementNode(userA))))
	comments := commentsA.FirstChild.Data
	fmt.Println(comments)
	fmt.Println(user)
	fmt.Println(score)
	return "", 0, 0, nil
}

func findPosts(node *html.Node) *html.Node {
	table := tagFinder(node, tableTag, postsTableAttrVal)
	if table == nil {
		return nil
	}

	tb := getFirstChildElementNode(table)
	if tb.Data != tbodyTag {
		return nil
	}
	return tb
}

//getNextElement
//goToElement
//getElementAttributes
//getElementText
