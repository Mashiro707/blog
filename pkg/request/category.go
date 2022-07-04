package request

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateCategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DeleteCategoryRequest struct {
	ID int `json:"id"`
}
