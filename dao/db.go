package dao

import (
	"errors"
	"log"
	"project1/http_models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //声明全局变量（数据库)(入口)
// InitDB 连接数据库
func InitDB(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	//自动迁移表结构
	err = db.AutoMigrate(&http_models.User{})
	if err != nil {
		return err
	}
	DB = db
	log.Println("数据库连接成功!")
	return nil
}

// CreateUser 创建用户
func CreateUser(user http_models.User) error {
	//密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return DB.Create(&user).Error
}

// UserLoginByPhone 手机号登录
func UserLoginByPhone(phone, password string) (http_models.User, error) {
	user := http_models.User{}
	err := DB.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http_models.User{}, errors.New("用户不存在")
		}
		return http_models.User{}, err
	}
	//验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return http_models.User{}, errors.New("密码错误")
	}
	return user, nil
}

// UserLoginByEmail 邮箱登录
func UserLoginByEmail(email, password string) (http_models.User, error) {
	user := http_models.User{}
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http_models.User{}, errors.New("用户不存在")
		}
		return http_models.User{}, err
	}
	//验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return http_models.User{}, errors.New("密码错误")
	}
	return user, nil
}

// GetUserByID 根据ID获取用户
func GetUserByID(userID uint) (http_models.User, error) {
	user := http_models.User{}
	err := DB.Where("user_id = ?", userID).First(&user).Error // 改为按 user_id 查询
	return user, err
}

// CheckUserExists 检查用户是否已存在
func CheckUserExists(username, email, phone string) (bool, error) {
	var count int64
	err := DB.Model(&http_models.User{}).Where("username = ? OR email = ? OR phone = ?", username, email, phone).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
