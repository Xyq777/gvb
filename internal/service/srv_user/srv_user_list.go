package srv_user

import (
	"gvb/internal/models/ctype"
	dao2 "gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/tools/desense"
)

func (s UserSrv) UserList(page req.Page, role ctype.Role) (*res.Response, error) {
	users, count, err := dao2.GetList(dao2.UserModel{}, &page)
	if err != nil {
		resp := res.NewResponse(res.DatabaseOperateError, res.EmptyData, res.CodeMsg(res.DatabaseOperateError))
		return resp, err
	}
	var userList []dao2.UserModel
	for _, user := range users {
		if role != ctype.PermissionAdmin {
			// 脱敏
			user.Username = ""
			user.Tel = desense.TelDesensitization(user.Tel)
			user.Email = desense.EmailDesensitization(user.Email)
		}
		userList = append(userList, user)
	}
	listData := res.ListData{
		Count:     count,
		ModelList: userList,
	}
	resp := res.NewResponse(res.Success, listData, res.CodeMsg(res.Success))

	return resp, nil
}
