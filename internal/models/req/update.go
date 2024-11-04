package req

type UpdateReq struct {
	ID uint `json:"id" form:"id" binding:"required"`
}
type UpdateImageNameReq struct {
	UpdateReq
	ImageName string `json:"image_name" form:"image_name" binding:"required"`
}

func (r UpdateReq) GetIDList() []uint {
	return []uint{r.ID}
}
