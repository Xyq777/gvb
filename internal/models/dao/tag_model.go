package dao

import "gorm.io/gorm"

// TagModel 标签表
type TagModel struct {
	gorm.Model
	Title string `gorm:"size:16;comment:标签的名称;unique" json:"title"` // 标签的名称
}

func (t TagModel) FindWithTitle(tx *gorm.DB) (bool, error) {
	result := tx.Find(&t).Where("title = ?", t.Title)
	if result.RowsAffected == 0 {
		return false, result.Error
	}
	return true, result.Error
}
func (t TagModel) Create(tx *gorm.DB) error {
	return tx.Create(t).Error
}
func (t TagModel) Update(tx *gorm.DB) error {
	return tx.Updates(t).Where("tittle = ?", t.Title).Error
}
func (t TagModel) Delete(tx *gorm.DB) error {
	return tx.Model(t).Where("tittle = ?", t.Title).Delete(&t).Error
}
