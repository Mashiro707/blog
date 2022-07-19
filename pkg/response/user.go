package response

type UserInfo struct {
	ID          int64  `json:"id"`
	UserName    string `json:"user_name"`
	NickName    string `json:"nick_name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	CreatedTime int64  `json:"created_time"`
	UpdatedTime int64  `json:"updated_time"`
}

type UserInfoWithToken struct {
	UserInfo
	Token string `json:"token"`
}
