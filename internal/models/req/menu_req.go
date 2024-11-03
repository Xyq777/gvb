package req

import "gvb/internal/models/ctype"

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	MenuTitle     string      `json:"menu_title" binding:"required" msg:"请完善菜单名称"`
	MenuTitleEn   string      `json:"menu_title_en" binding:"required" msg:"请完善菜单英文名称"`
	Slogan        string      `json:"slogan"`
	Abstract      ctype.Array `json:"abstract"`
	AbstractTime  int         `json:"abstract_time"`                         // 切换的时间，单位秒
	BannerTime    int         `json:"banner_time"`                           // 切换的时间，单位秒
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号"` // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list"`                       // 具体图片的顺序
}
