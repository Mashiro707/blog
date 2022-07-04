package request

type CreateTagRequest struct {
	Name string `json:"name"`
}

type UpdateTagRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DeleteTagRequest struct {
	IDs []int `json:"ids"`
}
