package req

type Page struct {
	Page  int    `form:"page" binding:"required"`
	Key   string `form:"key" binding:"required"`
	Limit int    `form:"limit" binding:"required"`
	Sort  string `form:"sort" `
}
