package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` // 分类表的唯一标识
	Name     string `gorm:"column:name;type:varchar(100);" json:"name"`        // 分类名称
	ParentId int    `gorm:"column:parentId;type:int(11);" json:"parentId"`     // 父级ID
}

func (table *Category) TableName() string {
	return "category"
}
