package request

type CreateCategoryRequest struct {
	Name string `json:"name" form:"name"`
}

type UpdateCategoryRequest struct {
	ID   int64  `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

type DeleteCategoryRequest struct {
	ID int64 `json:"id" form:"id"`
}
