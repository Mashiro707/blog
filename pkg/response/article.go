package response

type ArticleList struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Cover        string `json:"cover"`
	IsTop        int8   `json:"is_top"`
	IsComment    int8   `json:"is_comment"`
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	TimeModel
}

type ArticleDetail struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Cover        string `json:"cover"`
	Content      string `json:"content"`
	Description  string `json:"description"`
	Status       int8   `json:"status"`
	IsTop        int8   `json:"is_top"`
	IsComment    int8   `json:"is_comment"`
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
