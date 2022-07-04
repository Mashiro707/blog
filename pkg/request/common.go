package request

// Pagination 分页
type Pagination struct {
	PageNum  int `json:"page_num" form:"page_num" `                          // 页码
	PageSize int `json:"page_size" form:"page_size" binding:"min=1,max=100"` // 分页数
}
