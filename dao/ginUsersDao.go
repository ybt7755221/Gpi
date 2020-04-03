package dao

import "time"

type GinUsersDao struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Fullname   string    `json:"fullname"xorm:"not null default '' comment('用户中文名') VARCHAR(50)"`
	Mobile     string    `json:"mobile" xorm:"not null default '' comment('手机号') unique CHAR(20)"`
	Email      string    `json:"email" xorm:"not null default 'example@example.com' comment('邮箱') VARCHAR(128)"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	UpdateTime time.Time `json:"update_time" xorm:"updated"`
}
