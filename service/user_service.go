package service

import (
	"Supply/Supply_and_Demand/dao"
	http_models2 "Supply/Supply_and_Demand/http_models"
	"errors"
	"fmt"
	"project1/utils"
	"time"
)

// UserRegister 用户注册服务(前后端交互)
func UserRegister(registerInfo http_models2.UserRegister) (uint, error) {
	// 检查用户名、邮箱、手机号是否已存在
	exists, err := dao.CheckUserExists(registerInfo.Username, registerInfo.Email, registerInfo.Phone)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, errors.New("用户名、邮箱或者手机号已存在")
	}

	//数据转换(DTO->实体)
	user := http_models2.User{
		Username: registerInfo.Username,
		Password: registerInfo.Password,
		Email:    registerInfo.Email,
		Phone:    registerInfo.Phone,
	}
	err = dao.CreateUser(user)
	if err != nil {
		fmt.Println("注册失败:", err.Error())
		return 0, err
	}

	// 重新查询获取用户ID（因为CreateUser中的user可能没有返回ID）
	newUser := http_models2.User{}
	err = dao.DB.Where("phone = ?", registerInfo.Phone).First(&newUser).Error
	if err != nil {
		return 0, err
	}

	return newUser.UserID, nil
}

// UserLoginByPhone 手机号登录服务
func UserLoginByPhone(LoginInfo *http_models2.UserLoginByPhone) (*http_models2.LoginSuccessData, error) {
	//身份校验
	user, err := dao.UserLoginByPhone(LoginInfo.Phone, LoginInfo.Password)
	if err != nil {
		return nil, err
	}
	//生成jwt token
	token, err := utils.GenerateToken(user.UserID, user.Username)
	if err != nil {
		return nil, err
	}
	//更新最后登录时间
	dao.DB.Model(&user).Update("last_login", time.Now())
	return &http_models2.LoginSuccessData{ //创建结构体指针
		UserID:   user.UserID,
		Username: user.Username,
		Token:    token,
		Expire:   time.Now().Add(7 * 24 * time.Hour).Unix(),
	}, nil
}

// UserLoginByEmail 邮箱登录
func UserLoginByEmail(loginInfo *http_models2.UserLoginByEmail) (*http_models2.LoginSuccessData, error) {
	user, err := dao.UserLoginByEmail(loginInfo.Email, loginInfo.Password)
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateToken(user.UserID, user.Username)
	if err != nil {
		return nil, err
	}
	//更新最后登录时间
	dao.DB.Model(&user).Update("last_login", time.Now())
	return &http_models2.LoginSuccessData{ //创建结构体指针
		UserID:   user.UserID,
		Username: user.Username,
		Token:    token,
		Expire:   time.Now().Add(7 * 24 * time.Hour).Unix(),
	}, nil

}
