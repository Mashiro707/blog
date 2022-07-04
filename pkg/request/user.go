package request

type UserLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
