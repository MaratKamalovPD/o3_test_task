package args

type CreateCommentArgs struct {
	PostID          int    `json:"postId"`
	ParentCommentID *uint  `json:"parentCommentId"`
	UserID          int    `json:"userId"`
	Content         string `json:"content"`
}

type GetCommentsArgs struct {
	PostID          int   `json:"postId"`
	ParentCommentID *uint `json:"parentCommentId"`
	Limit           int   `json:"limit"`
	Offset          int   `json:"offset"`
}
