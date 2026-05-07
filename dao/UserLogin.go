package dao

import (
	"Supply_and_Demand/config"
	"Supply_and_Demand/global"
	"Supply_and_Demand/http_models"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// redis 手动赋值
func init() {
	if config.RedisClient == nil {
		fmt.Println("正在初始化 Redis 客户端")
		config.RedisClient = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := config.RedisClient.Ping(ctx).Result()
		if err != nil {
			fmt.Printf("Redis 连接失败：%v\n", err)
			config.RedisClient = nil
		} else {
			fmt.Println("Redis 客户端初始化成功！")
		}
	}
}

type UserLoginDao struct {
	Orm *gorm.DB
}

func NewUserLogin() *UserLoginDao {
	return &UserLoginDao{Orm: global.DB}
}

func (m *UserLoginDao) CheckUserByEmail(email string) (bool, error) {
	UserLogin := http_models.UserLogin{}
	err := m.Orm.Where("email = ?", email).First(&UserLogin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *UserLoginDao) CheckUserByPhoneNumber(phoneNumber string) (bool, error) {
	UserLogin := http_models.UserLogin{}
	err := m.Orm.Where("phone_number = ?", phoneNumber).First(&UserLogin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (m *UserLoginDao) CreateUser(UserLogin http_models.UserLogin) error {
	result := m.Orm.Create(&UserLogin)
	if result.Error != nil {
		return errors.New("用户创建失败: " + result.Error.Error())
	}
	return nil
}
func (m *UserLoginDao) LoginUserByPhoneNumber(PhoneNumber string) (*http_models.UserLogin, error) {
	UserLogin := &http_models.UserLogin{}
	err := m.Orm.Where("phone_number = ?", PhoneNumber).First(UserLogin).Error
	if err != nil {
		return nil, err
	}
	return UserLogin, nil
}
func (m *UserLoginDao) LoginUserByEmail(Email string) (*http_models.UserLogin, error) {
	UserLogin := &http_models.UserLogin{}
	err := m.Orm.Where("email = ?", Email).First(UserLogin).Error
	if err != nil {
		return nil, err
	}
	return UserLogin, nil
}
func (m *UserLoginDao) GetUserByEmail(email string) (*http_models.UserLogin, error) {
	UserLogin := &http_models.UserLogin{}
	err := m.Orm.Where("email = ?", email).First(UserLogin).Error
	if err != nil {
		return nil, err
	}
	return UserLogin, nil
}

func (m *UserLoginDao) GetUserByTelephoneNumber(telephoneNumber string) error {
	UserLogin := http_models.UserLogin{}
	err := m.Orm.Where("phone_number= ?", telephoneNumber).Find(&UserLogin).Error
	if err != nil {
		return err
	}
	return nil

}
func (m *UserLoginDao) GetUserByID(ID uint) error {
	UserLogin := http_models.UserLogin{}
	err := m.Orm.Where("ID= ?", ID).Find(&UserLogin).Error
	if err != nil {
		return err
	}
	return nil
}
func (m *UserLoginDao) UpdateUser(UserLogin *http_models.UserLogin) error {
	result := m.Orm.Save(UserLogin)
	if result.Error != nil {
		return errors.New("用户更新失败" + result.Error.Error())
	}
	return nil
}
func (m *UserLoginDao) UpdateUserName(ID uint, UserName string) error {
	UserLogin := &http_models.UserLogin{}
	err := m.Orm.Where("ID= ?", ID).Find(&UserLogin).Error
	if err != nil {
		return errors.New("查找用户失败")
	}
	UserLogin.UserName = UserName
	UserLogin.UpdateTime = time.Now()
	return m.UpdateUser(UserLogin)
}
func (m *UserLoginDao) UpdateUserAvatar(ID uint, Avatar string) error {
	UserLogin := &http_models.UserLogin{}
	err := m.Orm.Where("id= ?", ID).Find(&UserLogin).Error
	if err != nil {
		return errors.New("未查询到用户")
	}
	UserLogin.Avatar = Avatar
	UserLogin.UpdateTime = time.Now()
	return m.UpdateUser(UserLogin)
}

// SaveSmsCode 保存验证码到Redis
func (m *UserLoginDao) SaveSmsCode(phoneNumber, code string, expiration time.Duration) error {
	ctx := context.Background()
	key := fmt.Sprintf("sms_code:%s", phoneNumber)
	err := config.RedisClient.Set(ctx, key, code, expiration).Err()
	if err != nil {
		return errors.New("验证码保存失败")
	}
	return nil
}

// GetSmsCode 从redis里面获取验证码
func (m *UserLoginDao) GetSmsCode(phoneNumber string) (string, error) {
	ctx := context.Background()
	key := fmt.Sprintf("sms_code:%s", phoneNumber)
	code, err := config.RedisClient.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", errors.New("验证码过期或者不存在了")
		}
		return "", errors.New("获取验证码失败")
	}
	return code, nil
}

// DeleteSmsCode 删除验证码
func (m *UserLoginDao) DeleteSmsCode(phoneNumber string) error {
	ctx := context.Background()
	key := fmt.Sprintf("sms_code:%s", phoneNumber)
	err := config.RedisClient.Del(ctx, key).Err()
	if err != nil {
		return errors.New("删除失败")
	}
	return nil
}

// CheckSmsCodeSendLimit 检查频率发送机制(60s)
func (m *UserLoginDao) CheckSmsCodeSendLimit(phoneNumber string) (bool, error) {
	ctx := context.Background()
	limitKey := fmt.Sprintf("sms_limit:%s", phoneNumber)
	exists, err := config.RedisClient.Exists(ctx, limitKey).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

// SetSmsCodeLimit 设置发送频率
func (m *UserLoginDao) SetSmsCodeLimit(phoneNumber string, expiration time.Duration) error {
	ctx := context.Background()
	LimitKey := fmt.Sprintf("sms_limit:%s", phoneNumber)
	err := config.RedisClient.Set(ctx, LimitKey, "1", expiration).Err()
	if err != nil {
		return errors.New("设置发送限制失败")
	}
	return nil
}
