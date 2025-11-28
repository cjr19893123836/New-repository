package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your-secret-key-change-in-production") // 生产环境要更改,jwt密钥
// 声明jwt结构体
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 创建令牌
func GenerateToken(userID uint, username string) (string, error) {
	//创建声明结构体
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), //过期时间7天
			IssuedAt:  jwt.NewNumericDate(time.Now()),                         //签发时间
			Issuer:    "auth-app",                                             //签发这(初始化模块名)
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //创建令牌(哈希)
	return token.SignedString(jwtSecret)                       //进行签名并获取完整令牌
}

// ParseToken 解析令牌
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	//类型断言
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil //有效的令牌
	}
	return nil, jwt.ErrSignatureInvalid //无效的令牌
}
