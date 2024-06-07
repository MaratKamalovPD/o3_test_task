package args

type PostArgs struct {
	ID int `json:"id"`
}

type CreatePostArgs struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"userId"`
}

type DisableCommentsArgs struct {
	PostID int `json:"postId"`
	UserID int `json:"userId"`
}
