package main

import (
	"fmt"
	"log"
	"math/rand"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"context"
	"net/http"
	"regexp"
	"time"
)
var ctx = context.Background()
var rdb *redis.Client
func RandomCode() string {
	return fmt.Sprintf("%06d", rand.Intn(100000))
}
func isphonenuminvalid(phone string) bool {
	reg:=regexp.MustCompile("^1[3-9]/d[9]$")
	return !reg.MatchString(phone)
}
func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		PoolSize: 10000, // 连接池最大连接数
	})
	if err:=rdb.Ping(ctx).Err();
	err!=nil{
		panic(fmt.Sprintf("connect redis failed, err:%v",err))
	}
}
func RandID() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s:=make([]rune, n)
	for i:=range s{
		s[i]=letters[rand.Intn(len(letters))]
	}
	return string(s)
}
type user struct {
	Phone string 'json:"phone" gorm:"phone"`'
	Nickname string 'json:"nickname"gorm:"nickname"`'
}
func SendCodeHandler(c *gin.Context) {
	phone:=c.PostForm("phone")
	a:=isphonenuminvalid(phone)
	if a==true{
		c.String(http.StatusOK,"手机格式错误")
		return
}
	code:=RandomCode()
	err := rdb.Set(ctx, "login:code:"+phone,code,5*time.Minute).Err()
	if err != nil {
		c.String(http.StatusOK,"获取验证码失败")
		return
	}
	log.Printf("验证码为:%s",code)
	c.String(200,"验证码发送成功" )
	//还有那个用运营商发送验证码未完成
}
	func LoginHandler(c *gin.Context) {
		var form struct {
			Phone string `json:"phone"`
			Code  string `json:"code"`
		}
		 err:=c.ShouldBind(&form)
		 if err!=nil{
			 c.JSON(200,gin.H{"error":"参数错误" })
			return
		 }
		 if a:=isphonenuminvalid(form.Phone); a==true{
			 c.JSON(200,gin.H{"error":"手机号格式错误"})
			 return
		 }
		 val,err:=rdb.Get(ctx, "login:code"+form.Phone).Result(){
			 if err==redis.Nil{
					c.JSON(200,gin.H{"error":"验证码已经过期"})
					return
			 }
			 if val!=form.Code{
				 c.JSON(200,gin.H{"error":"验证码错误,请重新输入"})
			 }
			rdb.Del(ctx, "login:code:"+form.Phone) //删除标签

		}

	}
func main (){

}