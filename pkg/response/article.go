package response

type ArticleList struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Description  string `json:"description"`
	IsTop        bool   `json:"is_top"`
	IsComment    bool   `json:"is_comment"`
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
}

type ArticleDetail struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Description  string `json:"description"`
	IsTop        bool   `json:"is_top"`
	IsComment    bool   `json:"is_comment"`
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
}
