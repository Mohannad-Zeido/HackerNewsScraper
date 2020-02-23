package types

const (
	TableTag                     = "table"
	GeneralInfoTag               = "athing"
	PostsTableAttrVal            = "itemlist"
	TbodyTag                     = "tbody"
	URIAttr                      = "href"
	DetailsTag                   = "subtext"
	EndOfPostsAttrVal            = "morespace"
	NonNumbers                   = "[^\\d|\\-]"
	InternalURI                  = "^item\\?id=[0-9a-zA-Z]*"
	NumberOfNodesPerPost         = 3
	URINodePositionInGeneralInfo = 2
	RankNodeDepth                = 2
	CommentsNodePosition         = 5
)

//Post is the struct used for storing a post
type Post struct {
	Title    string `json:"title"`
	URI      string `json:"uri"`
	Author   string `json:"author"`
	Points   int    `json:"points"`
	Comments int    `json:"comments"`
	Rank     int    `json:"rank"`
}

//RunState is a struct used for indicating the run state (either test or normal ie. production)
type RunState struct {
	TestFile  string
	TestState bool
}
