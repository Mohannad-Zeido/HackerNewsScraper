package helper

import "golang.org/x/net/html"

func TagFinder(node *html.Node, tagName, attrValue string) *html.Node {
	if node == nil {
		return nil
	}

	if node.Type == html.ElementNode && node.Data == tagName && ContainsAttributeValue(node.Attr, attrValue) {
		return node
	}

	for child := GetFirstChildElement(node); child != nil; child = GetNextSiblingElement(child) {
		result := TagFinder(child, tagName, attrValue)
		if result != nil {
			return result
		}
	}
	return nil
}

func ContainsAttributeValue(attributes []html.Attribute, attributeValue string) bool {
	for _, a := range attributes {
		if a.Val == attributeValue {
			return true
		}
	}
	return false
}

func AttributeValue(attributes []html.Attribute, attribute string) string {
	for _, a := range attributes {
		if a.Key == attribute {
			return a.Val
		}
	}
	return ""
}

func GetTagText(node *html.Node) (string, bool) {
	textNode := node.FirstChild
	if textNode == nil {
		return "", false
	}
	return textNode.Data, true
}

func GetNextSiblingElement(node *html.Node) *html.Node {
	for sibling := node.NextSibling; sibling != nil; sibling = sibling.NextSibling {
		if sibling.Type == html.ElementNode {
			return sibling
		}
	}
	return nil
}

func GetFirstChildElement(node *html.Node) *html.Node {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode {
			return child
		}
	}
	return nil
}

func GetNthSibling(node *html.Node, n int) *html.Node {
	for i := 0; i < n; i++ {
		node = GetNextSiblingElement(node)
		if node == nil {
			return nil
		}
	}
	return node
}

func GetChildAtDepth(node *html.Node, n int) *html.Node {
	for i := 0; i < n; i++ {
		node = GetFirstChildElement(node)
		if node == nil {
			return nil
		}
	}
	return node
}
