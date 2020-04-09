package entities

import (
	"time"
)

type GinContents struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Topic      string    `json:"topic" xorm:"not null comment('主题') VARCHAR(255)"`
	Content    string    `json:"content" xorm:"not null comment('详细内容') TINYTEXT"`
	Category   int       `json:"category" xorm:"not null default 0 comment('分类') TINYINT(3)"`
	TestTime   time.Time `json:"test_time" xorm:"not null comment('测试时间') DATETIME"`
	PulishTime time.Time `json:"publish_time" xorm:"not null comment('上线时间') DATETIME"`
	OpTime     time.Time `json:"op_time" xorm:"updated"`
}

type GinContentsPageDao struct {
	List     []GinContents `json:"list"`
	PageNum  int           `json:"page_num"`
	PageSize int           `json:"page_size"`
	Total    int64         `json:"total"`
}
