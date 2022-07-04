package request

type CreateArticle struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Content     string `json:"content"`
	Description string `json:"description"`
	IsTop       int8   `json:"is_top"`
	IsComment   int8   `json:"is_comment"`
	CategoryID  int    `json:"category_id"`
	Tags        []int  `json:"tags"`
}

type UpdateArticle struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Content     string `json:"content"`
	Description string `json:"description"`
	IsTop       int8   `json:"is_top"`
	IsComment   int8   `json:"is_comment"`
	CategoryID  int    `json:"category_id"`
	Tags        []int  `json:"tags"`
}

type DeleteArticle struct {
	ID int `json:"id"`
}

type GetArticleByID struct {
	ID int `json:"id"`
}
