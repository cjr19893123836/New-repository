package http_models

import "gorm.io/gorm"

type Homepage struct {
	gorm.Model
	UserID        int64  `gorm:"column:user_id;unique;not null"`   //用户ID
	BackgroundURL string `gorm:"column:background_url;default:''"` //背景图片链接
	AvatarBanner  string `gorm:"column:avatar_banner;default:''"`  //头像图片
	Signature     string `gorm:"column:signature;default:''"`      //个性签名
	PersonalLabel string `gorm:"column:personal_label;default:''"` //个人标签
	IsPublic      int    `gorm:"column:is_public;default:1"`       //主页公开状态
	ShowFans      int    `gorm:"column:show_fans;default:1"`       //粉丝数展开开关
	ShowFollow    int    `gorm:"column:show_follow;default:1"`     //关注数展开开关
	LayoutContent string `gorm:"column:layout_content"`            //主页布局
}
