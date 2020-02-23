package types

import (
	"regexp"
)

const (
	TableTag                       = "table"
	GeneralInfoTag                 = "athing"
	PostsTableAttrVal              = "itemlist"
	TbodyTag                       = "tbody"
	UriAttr                        = "href"
	DetailsTag                     = "subtext"
	EndOfPostsAttrVal              = "morespace"
	nonNumbers                     = "[^\\d|\\-]"
	internalUri                    = "^item\\?id=[0-9a-zA-Z]*"
	NumberOfNodesPerPost           = 3
	UriNodePositionInGeneralInfo   = 2
	TableNodePositionInGeneralInfo = 3
	RankNodeDepth                  = 2
	CommentsNodePosition           = 5
)

var (
	InternalUriRegex, _ = regexp.Compile(internalUri)
	NonNumbersRegex, _  = regexp.Compile(nonNumbers)
)

type Post struct {
	Title    string `json:"title"`
	Uri      string `json:"uri"`
	Author   string `json:"author"`
	Points   int    `json:"points"`
	Comments int    `json:"comments"`
	Rank     int    `json:"rank"`
}

type RunState struct {
	TestFile  string
	TestState bool
}
