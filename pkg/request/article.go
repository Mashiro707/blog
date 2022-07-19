package request

type CreateArticle struct {
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Content     string `json:"content"`
	Description string `json:"description"`
	IsTop       int8   `json:"is_top"`
	IsComment   int8   `json:"is_comment"`
	CategoryID  int64  `json:"category_id"`
}

type UpdateArticle struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Content     string `json:"content"`
	Description string `json:"description"`
	IsTop       int8   `json:"is_top"`
	IsComment   int8   `json:"is_comment"`
	CategoryID  int64  `json:"category_id"`
}

type DeleteArticle struct {
	ID int64 `json:"id"`
}

type GetArticleByID struct {
	ID int64 `json:"id"`
}
