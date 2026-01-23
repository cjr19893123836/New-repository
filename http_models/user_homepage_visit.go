package http_models

import (
	"time"

	"gorm.io/gorm"
)

type UserHomepageVisit struct {
	gorm.Model
	HomepageID     int           `gorm:"column:homepage_id"`     //主表ID
	VisitorUserID  int           `gorm:"column:visitor_user_id"` //访客用户id
	VisitorIP      string        `gorm:"column:visitor_ip"`      //访客用户id
	VisitorDevice  string        `gorm:"column:visitor_device"`  //访客设备信息
	VisitorTime    time.Time     `gorm:"column:visitor_time"`    //	访问时间
	StadyDurdation time.Duration `gorm:"column:stady_durdation"` //停留时长
	EntryPage      string        `gorm:"column:entry_page"`      //入口页面
}
