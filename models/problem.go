package models

import (
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`        // 问题表的唯一标识
	CategoryId string `gorm:"column:category_id;type:varchar(255);" json:"catogory_id"` // 分类ID,以逗号分隔
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`             // 问题标题
	Content    string `gorm:"column:content;type:text;" json:"content"`                 // 问题内容
	MaxRuntime int    `gorm:"column:max_runtime;type:text;" json:"max_runtime"`         // 最大运行时间
	MaxMem     int    `gorm:"column:max_mem;type:text;" json:"max_mem"`                 // 最大运行内存
}

func (table *Problem) TableName() string {
	return "problem"
}
