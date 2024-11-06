package srv_user

import (
	"gvb/internal/dao"
	"gvb/internal/models"
	"gvb/internal/models/ctype"
	"gvb/internal/models/serializition/req"
	"gvb/internal/models/serializition/res"
	"gvb/tools/desense"
)

func (s UserSrv) UserList(page req.Page, role ctype.Role) (*res.Response, error) {
	users, count, err := dao.GetList(models.UserModel{}, &page)
	if err != nil {
		resp := res.NewResponse(res.DatabaseOperateError, res.EmptyData, res.CodeMsg(res.DatabaseOperateError))
		return resp, err
	}
	var userList []models.UserModel
	for _, user := range users {
		if role != ctype.PermissionAdmin {
			// 脱敏
			user.Username = ""
			user.Tel = desense.TelDesensitization(user.Tel)
			user.Email = desense.EmailDesensitization(user.Email)
		}
		userList = append(userList, user)
	}
	listData := res.ListRespData{
		Count: count,
		List:  userList,
	}
	resp := res.NewResponse(res.Success, listData, res.CodeMsg(res.Success))

	return resp, nil
}
