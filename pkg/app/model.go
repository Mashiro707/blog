package app

// Response ...
type Response struct {
	Code    int         `json:"status_code" example:"200"`
	Msg     string      `json:"status_message" example:"ok"`
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data"`
}

// Page 分页
type Page struct {
	List     interface{} `json:"list"`
	Count    int64       `json:"count"`
	PageNo   int         `json:"page_no"`
	PageSize int         `json:"page_size"`
}

// PageResponse 分页response
type PageResponse struct {
	Code    int    `json:"status_code" example:"200"`
	Msg     string `json:"status_message" example:"ok"`
	Success bool   `json:"success" example:"true"`
	Data    Page   `json:"data"`
}

// ReturnOK 成功返回
func (res *Response) ReturnOK() *Response {
	res.Code = 200
	res.Success = true
	return res
}

// ReturnError 错误返回
func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	res.Success = false
	return res
}

// ReturnOK 分页成功返回
func (res *PageResponse) ReturnOK() *PageResponse {
	res.Code = 200
	res.Success = true
	return res
}

// ReturnError 错误返回
func (res *PageResponse) ReturnError(code int) *PageResponse {
	res.Code = code
	res.Success = false
	return res
}
