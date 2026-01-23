package http_models

import "time"

type UserHomepageTop struct {
	ID          int64     `gorm:"column:id"`             //置顶记录唯一ID
	HomepageId  int64     `gorm:"column:homepage_id"`    //主表ID
	ContentType int8      `gorm:"column:content_type;"`  //内容类型
	ContentId   int64     `gorm:"column:content_id"`     //对应类型内容ID
	TopSort     int8      `gorm:"column:top_sort"`       //置顶排序
	TopStartTim time.Time `gorm:"column:top_start_time"` //置顶开始时间
	TopEndTime  time.Time `gorm:"column:top_end_time"`   //置顶结束时间
	CreateTime  time.Time `gorm:"column:create_time"`    //记录创建时间
}
