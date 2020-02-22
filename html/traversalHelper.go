package html

import "golang.org/x/net/html"

func tagFinder(node *html.Node, tagName, attrValue string) *html.Node {
	if node == nil {
		return nil
	}

	if node.Type == html.ElementNode && node.Data == tagName && containsAttributeValue(node.Attr, attrValue) {
		return node
	}

	for child := getFirstChildElementNode(node); child != nil; child = getNextSiblingElementNode(child) {
		result := tagFinder(child, tagName, attrValue)
		if result != nil {
			return result
		}
	}
	return nil
}

func containsAttributeValue(attributes []html.Attribute, attributeValue string) bool {
	for _, a := range attributes {
		if a.Val == attributeValue {
			return true
		}
	}
	return false
}

func attributeValue(attributes []html.Attribute, attribute string) string {
	for _, a := range attributes {
		if a.Key == attribute {
			return a.Val
		}
	}
	return ""
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
