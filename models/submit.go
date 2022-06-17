package models

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(36);" json:"identity"`                 // 提交表的唯一标识
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"` // 问题表的唯一标识
	UserIdentity    string `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`       // 用户表的唯一标识
	Status          int    `gorm:"column:status;type:int(1);" json:"status"`                          //0:待判断 1：答案正确  2：答案错误  3：运行超内存 4：运行超时
}

func (table *Submit) TableName() string  {
	return "submit"
}
