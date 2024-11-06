package req

type UserEmailLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateReq struct {
	UserID   uint   `json:"user_id" binding:"required"`
	Nickname string `json:"Nickname" binding:"required"`
}
