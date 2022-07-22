package request

type CreateArticle struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	IsTop       bool   `json:"is_top"`
	IsComment   bool   `json:"is_comment"`
	CategoryID  int64  `json:"category_id"`
}

type UpdateArticle struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	IsTop       bool   `json:"is_top"`
	IsComment   bool   `json:"is_comment"`
	CategoryID  int64  `json:"category_id"`
}

type DeleteArticle struct {
	ID int64 `json:"id" form:"id"`
}

type GetArticleByID struct {
	ID int64 `json:"id" form:"id"`
}
