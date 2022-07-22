package app

// Response ...
type Response struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"message" example:"ok"`
	Data interface{} `json:"data"`
}

// Page 分页
type Page struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	PageNum  int         `json:"page_num"`
	PageSize int         `json:"page_size"`
}

// PageResponse 分页response
type PageResponse struct {
	Code int    `json:"code" example:"200"`
	Msg  string `json:"message" example:"ok"`
	Data Page   `json:"data"`
}

// ReturnOK 成功返回
func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

// ReturnError 错误返回
func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

// ReturnOK 分页成功返回
func (res *PageResponse) ReturnOK() *PageResponse {
	res.Code = 200
	return res
}

// ReturnError 分页错误返回
func (res *PageResponse) ReturnError(code int) *PageResponse {
	res.Code = code
	return res
}
