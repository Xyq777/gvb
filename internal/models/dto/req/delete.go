package req

type DeleteReq struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

type DeleteReqList []DeleteReq

func (l DeleteReqList) GetIDList() []uint {
	IDList := make([]uint, 0)
	for _, c := range l {
		IDList = append(IDList, c.ID)
	}
	return IDList
}
