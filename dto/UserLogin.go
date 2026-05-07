package dto

import (
	"Supply_and_Demand/http_models"
)

type LoginByEmailRep struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
type LoginByPhoneNumberRep struct {
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}
type RegisterRep struct {
	UserName    string `json:"user_name" binding:"required,min=3,max=20"`
	Password    string `json:"password" binding:"required,min=8,max=20"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"telephone" binding:"required,telephone"`
}

type LoginAndRegister struct {
	Token       string `json:"token" `
	ID          uint   `json:"id" `
	UserName    string `json:"username" `
	Email       string `json:"email" `
	PhoneNumber string `json:"telephone" `
}

func (rep *RegisterRep) ToModel() http_models.UserLogin {
	return http_models.UserLogin{
		Password:    rep.Password,
		Email:       rep.Email,
		UserName:    rep.UserName,
		PhoneNumber: rep.PhoneNumber,
	}
}
