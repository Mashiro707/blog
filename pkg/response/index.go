package response

type ArticleIndexList struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	IsTop        int8   `json:"is_top"`
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"categoryName"`
}
