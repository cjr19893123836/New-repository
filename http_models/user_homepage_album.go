package http_models

import "gorm.io/gorm"

type UserHomepageAlbum struct {
	gorm.Model
	HomepageID int64  `gorm:"column:homepage_id"`           //主表ID
	AlbumName  string `gorm:"column:album_name"`            //相册名称
	AlbumCover string `gorm:"column:album_cover"`           //相册封面图
	IsPublic   int    `gorm:"column:is_public;default:1"`   //相册公开状态
	SortOrder  int    `gorm:"column:sort_order ;default:1"` //相册排序
}
