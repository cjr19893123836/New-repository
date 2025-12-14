package http_models

import (
	"crypto/rand"
	"math/big"
	"time"

	"gorm.io/gorm"
)

// User 定义用户模型（结构体)
type User struct {
	gorm.Model
	UserID    uint      `gorm:"uniqueIndex;not null"`
	Username  string    `gorm:"uniqueIndex;size:50"`  //用户名，唯一索引
	Password  string    `gorm:"size:255"`             //密码，私密，不序列化到json
	Email     string    `gorm:"size:100;uniqueIndex"` //邮箱，唯一索引
	Phone     string    `gorm:"size:11"`
	LastLogin time.Time `gorm:"Default:NULL"` //最后登陆时间
}

// BeforeCreate 在创建用户前生成随机UserID
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.UserID == 0 {
		// 生成唯一的UserID
		userID, err := generateUniqueUserID(tx)
		if err != nil {
			return err
		}
		u.UserID = userID
	}
	return nil
}

// generateUniqueUserID 生成唯一的UserID
func generateUniqueUserID(tx *gorm.DB) (uint, error) {
	maxAttempts := 10 // 最大尝试次数

	for i := 0; i < maxAttempts; i++ {
		// 生成6-8位的随机数字 (100000 - 99999999)
		randomNum, err := rand.Int(rand.Reader, big.NewInt(98999999))
		if err != nil {
			return 0, err
		}

		userID := uint(randomNum.Int64() + 100000) // 确保最小6位数

		// 检查UserID是否已存在
		var count int64
		err = tx.Model(&User{}).Where("user_id = ?", userID).Count(&count).Error
		if err != nil {
			return 0, err
		}

		if count == 0 {
			return userID, nil // 找到唯一的UserID
		}
	}

	// 如果尝试多次都冲突，使用时间戳作为后备方案
	return uint(time.Now().Unix() % 100000000), nil
}
