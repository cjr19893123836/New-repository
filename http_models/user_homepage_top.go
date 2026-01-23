package http_models

import "time"

type UserHomepageAlbumImage struct {
	ID         int64     `gorm:"id"`          //相片唯一ID
	AlbumId    int64     `gorm:"album_id"`    //相册ID
	ImageUrl   string    `gorm:"image_url;"`  //图片URL
	ImageDesc  string    `gorm:"image_desc"`  //图片描述
	SortNum    int8      `gorm:"sort_num"`    //图片排序
	UploadTime time.Time `gorm:"upload_time"` //图片上传时间
}
