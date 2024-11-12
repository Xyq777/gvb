package req

type Page struct {
	Page  int    `form:"page" binding:"required" json:"page"`
	Key   string `form:"key" binding:"required" json:"key"`
	Limit int    `form:"limit" binding:"required" json:"limit"`
	Sort  string `form:"sort" json:"sort"`
}
