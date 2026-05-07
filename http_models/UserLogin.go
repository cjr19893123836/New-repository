package http_models

import "time"

type UserLogin struct {
	ID          uint       `json:"id" gorm:"column:id;primary_key"`
	UserName    string     `json:"user_name" gorm:"column:name;type:varchar(50);not null"`
	Password    string     `json:"password" gorm:"column:password;not null"`
	Email       string     `json:"email" gorm:"column:email;type:varchar(50);not null"`
	PhoneNumber string     `json:"phone_number" gorm:"column:phone_number;type:varchar(100);not null"`
	Avatar      string     `json:"avatar" gorm:"column:avatar;type:varchar(100);not null"`
	CreateTime  time.Time  `json:"create_time" gorm:"column:create_time;not null"`
	UpdateTime  time.Time  `json:"update_time" gorm:"column:update_time;not null"`
	LastLogin   *time.Time `json:"last_login" gorm:"column:last_login;default:NULL"`
}
