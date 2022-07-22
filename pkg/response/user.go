package response

type UserInfo struct {
	ID         int64  `json:"id"`
	UserName   string `json:"user_name"`
	NickName   string `json:"nick_name"`
	Email      string `json:"email"`
	Avatar     string `json:"avatar"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type UserInfoWithToken struct {
	UserInfo
	Token string `json:"token"`
}
