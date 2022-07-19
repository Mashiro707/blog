package response

type ArticleIndexList struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	IsTop        int8   `json:"is_top"`
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"categoryName"`
}
