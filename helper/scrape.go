package helper

//scrape.go provides helper functions for the html traversal

import "golang.org/x/net/html"

//TagFinder will find the HTML tag (identified with the tagName parameter) with the specified attribute value
//The attribute is not specified as all of them will be checked for the value
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

//ContainsAttributeValue will return true of the list of attributes contains the attributeValue and false if it does not
func ContainsAttributeValue(attributes []html.Attribute, attributeValue string) bool {
	for _, a := range attributes {
		if a.Val == attributeValue {
			return true
		}
	}
	return false
}

//AttributeValue will return the value assigned to the attribute
func AttributeValue(attributes []html.Attribute, attribute string) string {
	for _, a := range attributes {
		if a.Key == attribute {
			return a.Val
		}
	}
	return ""
}

//GetTagText will return the text in an HTML tag. If there is no text this function will return false
func GetTagText(node *html.Node) (string, bool) {
	textNode := node.FirstChild
	if textNode == nil {
		return "", false
	}
	return textNode.Data, true
}

//GetNextSiblingElement will return the nextSibling for the given HTML tag.
//this function skips textNodes as they are not HTML tags
func GetNextSiblingElement(node *html.Node) *html.Node {
	if node == nil {
		return nil
	}
	for sibling := node.NextSibling; sibling != nil; sibling = sibling.NextSibling {
		if sibling.Type == html.ElementNode {
			return sibling
		}
	}
	return nil
}

//GetFirstChildElement will return the fistChild for the given HTML tag.
//this function skips textNodes as they are not HTML tags
func GetFirstChildElement(node *html.Node) *html.Node {
	if node == nil {
		return nil
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode {
			return child
		}
	}
	return nil
}

//GetNthSibling will return the sibling of the HTML tag at a given positing
//this function will return nil if there are no more siblings
func GetNthSibling(node *html.Node, n int) *html.Node {
	for i := 0; i < n; i++ {
		node = GetNextSiblingElement(node)
		if node == nil {
			return nil
		}
	}
	return node
}

//GetNthSibling will return the child of the HTML tag at a given depth
//this function will return nil if there are no more sub-children
func GetChildAtDepth(node *html.Node, n int) *html.Node {
	for i := 0; i < n; i++ {
		node = GetFirstChildElement(node)
		if node == nil {
			return nil
		}
	}
	return node
}
