package req

type DeleteReq struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

func GetIDList(deleteReq []DeleteReq) []uint {
	var idList []uint
	for _, v := range deleteReq {
		idList = append(idList, v.ID)
	}
	return idList

}
