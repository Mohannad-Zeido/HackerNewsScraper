package html

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/parse"
	"golang.org/x/net/html"
)

type postNode struct {
	generalInfo *html.Node
	details     *html.Node
}

const (
	tableTag       = "table"
	generalInfoTag = "athing"
	postTableTag   = "itemlist"
	tbodyTag       = "tbody"
	detailsTag     = ""
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

		if child == nil || !containsAttribute(child.Attr, generalInfoTag) {
			return result
		}
		post.generalInfo = child

		child = getNextSiblingElementNode(child)
		if child == nil || len(child.Attr) != 0 {
			return result
		}
		post.details = child
		result = append(result, post)
		child = getNextSiblingElementNode(child)
		if child == nil {
			return result
		}
		child = getNextSiblingElementNode(child)
	}
}

func containsAttribute(attributes []html.Attribute, attribute string) bool {
	for _, a := range attributes {
		if a.Val == attribute {
			return true
		}
	}
	return false
}

func findPosts(node *html.Node) *html.Node {
	table := tableFinder(node)
	if table == nil {
		return nil
	}

	for child := getFirstChildElementNode(table); child != nil; child = getNextSiblingElementNode(child) {
		if child.Type == html.ElementNode && child.Data == tbodyTag {
			return child
		}
	}
	return nil
}

func tableFinder(node *html.Node) *html.Node {
	if node == nil {
		return nil
	}

	if node.Type == html.ElementNode && node.Data == tableTag && containsAttribute(node.Attr, postTableTag) {
		return node
	}

	for child := getFirstChildElementNode(node); child != nil; child = getNextSiblingElementNode(child) {
		result := tableFinder(child)
		if result != nil {
			return result
		}
	}
	return nil
}

func getNextSiblingElementNode(node *html.Node) *html.Node {
	for sibling := node.NextSibling; sibling != nil; sibling = sibling.NextSibling {
		if sibling.Type == html.ElementNode {
			return sibling
		}
	}
	return nil
}

func getFirstChildElementNode(node *html.Node) *html.Node {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode {
			return child
		}
	}
	return nil
}

//func getPostNode(node *html.Node) postNode{
//
//}

//getNextElement
//goToElement
//getElementAttributes
//getElementText
