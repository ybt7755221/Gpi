package entities

import (
	"time"
)

type WrUsers struct {
	Id       int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Username string    `json:"username" xorm:"not null default '' comment('用户名') VARCHAR(128)"`
	Mobile   string    `json:"mobile" xorm:"not null default '13000000000' comment('手机号') VARCHAR(20)"`
	Password string    `json:"password" xorm:"not null default '' comment('密码') VARCHAR(64)"`
	Email    string    `json:"email" xorm:"not null default '' comment('邮箱地址') unique VARCHAR(128)"`
	Created  time.Time `json:"created" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Updated  time.Time `json:"updated" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}

type WrUsersPageDao struct {
	List 	 []WrUsers	`json:"list"`
	PageNum  int		`json:"page_num"`
	PageSize int		`json:"page_size"`
	Total 	 int64		`json:"total"`
}
