package dto

import (
	"errors"
)

const (
	ERRConditionInvalid = errors.New("成色参数无效")
	ERRTradeTypeInvalid = errors.New("交易方式参数无效")
	ERRTransactionMethod= errors.New("邮寄方式参数无效")
)

type ProductReq struct {
	UserID      uint    `json:"user_id" binding:""`        //用户ID
	CategoryID   int64   `json:"category_id" binding:"required"`    //分类ID
	Title        string  `json:"title" binding:"required"`          //标题
	IntroText    string  `json:"intro_text" binding:"required"`     //描述
	OriginPrice  float64 `json:"origin_price" binding:"required"`   //原价（打折前）
	MainImageUrl string  `json:"main_image_url" binding:"required"` //主图
	Condition	string  `json:"condition" binding:"required"`      //成色
	TradeType    string  `json:"trade_type" binding:"required"`     //交易方式
	FreeShipping int     `json:"free_shipping" binding:"required"`  //邮寄方式
}

func (p *ProductReq) GetCondition() (string, error) {
	switch p.Condition {
	case "全新":
		return "全新", nil
	case "几乎全新":
		return "几乎全新", nil
	case "九成新":
		return "九成新", nil
	case "八成新":
		return "八成新", nil
	case "有破损"
		return "有破损", nil
	default:
		return "", ERRConditionInvalid
	}
}

func (p *ProductReq) GetTradeType() (string, error) {
	switch p.Condition {
	case "快递":
		return "快递", nil
	case "自提":
		return "自提", nil
	default:
		return "", ERRTradeTypeInvalidd
	}
}

func (p *ProductReq) GetTransactionMethod() (int, error) {
	switch p.Condition {
	// todo 补充交易方式


	default:
		return "", ERRTransactionMethod
	}
}