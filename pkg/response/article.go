package response

type ArticleList struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Status       int8   `json:"status"`
	IsTop        int8   `json:"is_top"`
	IsComment    int8   `json:"is_comment"`
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
	CreatedTime  int64  `json:"created_time"`
	UpdatedTime  int64  `json:"updated_time"`
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
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
}
