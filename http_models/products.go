package http_models

import (
	"Supply_and_Demand/dto"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID       uint      `gorm:"column:user_id"`        //用户ID
	CategoryID   uint      `gorm:"column:category_id"`    //分类ID
	Title        string    `gorm:"column:title"`          //标题
	IntroText    string    `gorm:"column:intro_text"`     //描述
	Price        float64   `gorm:"column:price"`          //价格
	OriginPrice  float64   `gorm:"column:origin_price"`   //原价（打折前）
	MainImageUrl string    `gorm:"column:main_image_url"` //主图
	PublishDate  time.Time `gorm:"column:publish_date"`   //发布时间
	Status       string    `gorm:"column:status"`         //商品状态
	ViewCount    int       `gorm:"column:view_count"`     //浏览量
	WantCount    int       `gorm:"column:want_count"`     //想要的数量
	Condition    string    `gorm:"column:condition"`      //成色
	TradeType    string    `gorm:"column:trade_type"`     //交易方式
	FreeShipping int       `gorm:"column:free_shipping"`  //邮寄方式
	CollectCount int       `gorm:"column:collect_count"`  //收藏量
	CreateDate   time.Time `gorm:"column:create_date"`    //商品发布时间
}

// 指定当前结构体对应的数据库表名
func (Product) TableName() string {
	return "products"
}

// ToResp 转换为DTO响应
func (p *Product) ToResp() dto.ProductResp {
	return dto.ProductResp{
		ID:           p.ID,
		UserID:       p.UserID,
		Title:        p.Title,
		IntroText:    p.IntroText,
		Price:        p.Price,
		OriginPrice:  p.OriginPrice,
		MainImageUrl: p.MainImageUrl,
		PublishDate:  p.PublishDate.Format("2006-01-02 15:04:05"),
		ViewCount:    p.ViewCount,
		WantCount:    p.WantCount,
		Condition:    p.Condition,
		TradeType:    p.TradeType,
		CollectCount: p.CollectCount,
	}
}
