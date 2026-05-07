package service

import (
	"Supply_and_Demand/dao"
	"Supply_and_Demand/http_models"
	"Supply_and_Demand/utils"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dypnsapi "github.com/alibabacloud-go/dypnsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type UserLoginService struct {
	Dao *dao.UserLoginDao
}

var userLogin *UserLoginService

func NewUserLoginService() *UserLoginService {
	if userLogin == nil {
		userLogin = &UserLoginService{
			Dao: dao.NewUserLogin(),
		}
	}
	return userLogin
}

// RegisterByEmail 注册
func (m *UserLoginService) RegisterByEmail(rep http_models.RegisterByEmail) (uint, error) {
	EmailExists, err := m.Dao.CheckUserByEmail(rep.Email)
	if err != nil {
		return 0, err
	}
	if EmailExists {
		return 0, errors.New("邮箱已经存在")
	}
	//验证码
	storedCode, err := m.Dao.GetSmsCode(rep.PhoneNumber)
	if err != nil {
		return 0, errors.New("验证码过期或者不存在")
	}
	if storedCode != rep.Code {
		return 0, errors.New("验证码错误")
	}
	//校验完后删除
	_ = m.Dao.DeleteSmsCode(rep.PhoneNumber)
	now := time.Now()
	UserNew := http_models.UserLogin{
		UserName:   rep.UserName,
		Password:   rep.Password,
		Email:      rep.Email,
		CreateTime: now,
		UpdateTime: now,
		LastLogin:  nil,
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserNew.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	UserNew.Password = string(hashedPassword)
	err = m.Dao.CreateUser(UserNew)
	if err != nil {
		return 0, err
	}
	return UserNew.ID, nil
}

func (m *UserLoginService) RegisterByPhone(rep http_models.RegisterByPhone) (uint, error) {
	PhoneNumberExists, err := m.Dao.CheckUserByPhoneNumber(rep.PhoneNumber)
	if err != nil {
		return 0, err
	}
	if PhoneNumberExists {
		return 0, errors.New("手机号已经存在")
	}
	// 从 Redis 获取验证码
	storedCode, err := m.Dao.GetSmsCode(rep.PhoneNumber)
	if err != nil {
		return 0, errors.New("验证码已过期或不存在，请重新获取")
	}
	if storedCode != rep.Code {
		return 0, errors.New("验证码错误")
	}
	// 校验完删除
	_ = m.Dao.DeleteSmsCode(rep.PhoneNumber)
	now := time.Now()
	UserNew := http_models.UserLogin{
		UserName:    rep.UserName,
		Password:    rep.Password,
		PhoneNumber: rep.PhoneNumber,
		CreateTime:  now,
		UpdateTime:  now,
		LastLogin:   nil,
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserNew.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	UserNew.Password = string(hashedPassword)

	err = m.Dao.CreateUser(UserNew)
	if err != nil {
		return 0, err
	}
	return UserNew.ID, nil
}

// LoginUserByPhone  登录
func (m *UserLoginService) LoginUserByPhone(rep http_models.LoginByPhoneNumberRep) (*http_models.LoginAndRegister, error) {
	User, err := m.Dao.LoginUserByPhoneNumber(rep.PhoneNumber)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	// 密码验证
	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(rep.Password))
	if err != nil {
		return nil, errors.New("密码错误")
	}
	now := time.Now()
	User.LastLogin = &now
	User.UpdateTime = now
	err = m.Dao.UpdateUser(User)
	if err != nil {
		fmt.Printf("更新用户登录时间失败：%v\n", err)
	}
	//校验验证码
	storedCode, err := m.Dao.GetSmsCode(rep.PhoneNumber)
	if err != nil {
		return nil, errors.New("验证码已过期或不存在，请重新获取")
	}
	if storedCode != rep.Code {
		return nil, errors.New("验证码错误")
	}
	//删除验证码
	_ = m.Dao.DeleteSmsCode(rep.PhoneNumber)
	//生成token
	token, err := utils.GenerateToken(User.ID, User.UserName)
	if err != nil {
		return nil, err
	}
	// 返回信息
	return &http_models.LoginAndRegister{
		ID:       User.ID,
		UserName: User.UserName,
		Token:    token,
		Expire:   time.Now().Add(time.Hour * 24).Unix(),
	}, nil
}

func (m *UserLoginService) LoginUserByEmail(rep http_models.LoginByEmailRep) (*http_models.LoginAndRegister, error) {
	User, err := m.Dao.LoginUserByEmail(rep.Email)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 密码验证
	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(rep.Password))
	if err != nil {
		return nil, errors.New("密码错误")
	}
	now := time.Now()
	User.LastLogin = &now
	User.UpdateTime = now
	err = m.Dao.UpdateUser(User)
	if err != nil {
		fmt.Printf("更新用户时间失败：%v\n", err)
	}
	//校验验证码
	storedCode, err := m.Dao.GetSmsCode(rep.PhoneNumber)
	if err != nil {
		return nil, errors.New("验证码已过期或不存在，请重新获取")
	}
	if storedCode != rep.Code {
		return nil, errors.New("验证码错误")
	}
	//删除验证码
	_ = m.Dao.DeleteSmsCode(rep.PhoneNumber)
	//生成token
	token, err := utils.GenerateToken(User.ID, User.UserName)
	if err != nil {
		return nil, err
	}
	// 返回信息
	return &http_models.LoginAndRegister{
		ID:       User.ID,
		UserName: User.UserName,
		Token:    token,
		Expire:   time.Now().Add(time.Hour * 24).Unix(),
	}, nil
}
func (m *UserLoginService) UpdateUserName(ID uint, UserName string) error {
	err := m.Dao.UpdateUserName(ID, UserName)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserAvatar 更新头像
func (m *UserLoginService) UpdateUserAvatar(ID uint, file *multipart.FileHeader) (string, error) {
	//提取文件的扩展名
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext == "" {
		ext = ".jpg"
	}
	//生成文件名
	filename := fmt.Sprintf("%d_%d%s", ID, time.Now().Unix(), ext)
	//创建一个目录
	uploadDir := "./upload/avatar/"
	if err := os.MkdirAll(uploadDir, 0777); err != nil {
		return "", err
	}
	//打开上传的文件
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func() { _ = src.Close() }()
	//创建目标文件
	dst, err := os.Create(filepath.Join(uploadDir, filename))
	if err != nil {

		return "", err
	}
	defer func() { _ = dst.Close() }()
	//复制文件内容
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}
	// 生成访问URL
	avatarUrl := "/upload/avatar/" + filename
	// 保存头像URL到数据库
	err = m.Dao.UpdateUserAvatar(ID, avatarUrl)
	if err != nil {
		// 如果数据库保存失败，删除已上传的文件
		_ = os.Remove(filepath.Join(uploadDir, filename))
		return "", err
	}

	return avatarUrl, nil
}

// 生成六位数的验证码
func generateSmsCode() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := 100000 + rng.Intn(900000)
	return fmt.Sprintf("%06d", code)
}

// sendSmsVerifyCodeByAliyun 阿里云短信发送验证码
func sendSmsVerifyCodeByAliyun(phoneNumber string, code string) error {
	// 创建配置
	config := &openapi.Config{
		AccessKeyId:     tea.String("LTAI5tA86XKwCm1CYGnb3S1S"),
		AccessKeySecret: tea.String("tUcbC0UQjPq175eIll1Ga15tRtsCqI"),
		Endpoint:        tea.String("dypnsapi.aliyuncs.com"),
	}

	client, err := dypnsapi.NewClient(config)
	if err != nil {
		return fmt.Errorf("创建阿里云客户端失败: %w", err)
	}
	// 构建请求参数
	templateParam := fmt.Sprintf(`{"code":"%s","min":"5"}`, code)
	request := &dypnsapi.SendSmsVerifyCodeRequest{
		PhoneNumber:   tea.String(phoneNumber),
		SignName:      tea.String("速通互联验证码"),
		TemplateCode:  tea.String("100001"),
		TemplateParam: tea.String(templateParam),
		CodeLength:    tea.Int64(6),
		ValidTime:     tea.Int64(300),
	}
	// 发送请求
	runtime := &util.RuntimeOptions{}
	resp, err := client.SendSmsVerifyCodeWithOptions(request, runtime)
	if err != nil {
		return fmt.Errorf("调用短信认证接口失败: %w", err)
	}
	fmt.Printf("阿里云短信认证响应: %+v\n", resp)

	if resp.Body.Code != nil && *resp.Body.Code != "OK" {
		codeStr := *resp.Body.Code
		message := ""
		if resp.Body.Message != nil {
			message = *resp.Body.Message
		}
		return fmt.Errorf("短信发送失败: %s - %s", codeStr, message)
	}

	return nil
}

// SendSmsCode 发送短信验证码
func (m *UserLoginService) SendSmsCode(phoneNumber string) error {
	// 检查发送频率限制
	limitExceeded, err := m.Dao.CheckSmsCodeSendLimit(phoneNumber)
	if err != nil {
		return err
	}
	if limitExceeded {
		return errors.New("请60秒之后重试")
	}
	// 生成6位验证码
	code := generateSmsCode()
	fmt.Printf("手机号：%s\n", phoneNumber)
	fmt.Printf("验证码：%s\n", code)
	// 保存验证码到Redis
	err = m.Dao.SaveSmsCode(phoneNumber, code, 5*time.Minute)
	if err != nil {
		return err
	}
	// 调用阿里云发送短信
	err = sendSmsVerifyCodeByAliyun(phoneNumber, code)
	if err != nil {
		_ = m.Dao.DeleteSmsCode(phoneNumber)
		return fmt.Errorf("短信发送失败: %w", err)
	}
	// 设置发送频率限制（60秒）
	err = m.Dao.SetSmsCodeLimit(phoneNumber, 60*time.Second)
	if err != nil {
		return err
	}
	return nil
}

// RegisterBySms 短信验证码注册
func (m *UserLoginService) RegisterBySms(rep http_models.RegisterBySmsReq) (*http_models.LoginAndRegister, error) {
	//检查用户是否存在
	exists, err := m.Dao.CheckUserByPhoneNumber(rep.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("手机号已经注册")
	}
	//从redis获得之前存的验证码
	code, err := m.Dao.GetSmsCode(rep.PhoneNumber)
	if err != nil {
		return nil, err
	}
	//对比验证码
	if code != rep.Code {
		return nil, errors.New("验证码错误")
	}
	//删除验证码
	_ = m.Dao.DeleteSmsCode(rep.PhoneNumber)
	//创建用户对象
	now := time.Now()
	UserNew := http_models.UserLogin{
		UserName:    rep.UserName,
		PhoneNumber: rep.PhoneNumber,
		CreateTime:  now,
		UpdateTime:  now,
		LastLogin:   nil,
		Password:    rep.Password,
	}
	//密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserNew.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	UserNew.Password = string(hashedPassword)
	//保存到数据库
	err = m.Dao.CreateUser(UserNew)
	if err != nil {
		return nil, err
	}
	//更新token
	token, err := utils.GenerateToken(UserNew.ID, UserNew.UserName)
	if err != nil {
		return nil, err
	}
	//返回Token和用户信息
	return &http_models.LoginAndRegister{
		ID:       UserNew.ID,
		UserName: UserNew.UserName,
		Token:    token,
		Avatar:   UserNew.Avatar,
		Expire:   time.Now().Add(time.Hour * 24).Unix(),
	}, nil
}

// LoginBySmsCode 短信验证码登录
func (m *UserLoginService) LoginBySmsCode(rep http_models.LoginBySmsRep) (*http_models.LoginAndRegister, error) {
	//验证短信验证码
	code, err := m.Dao.GetSmsCode(rep.PhoneNumber)
	if err != nil {
		return nil, err
	}
	//对比验证码是否正确
	if code != rep.Code {
		return nil, errors.New("验证码不正确")
	}
	//删除Redis的验证码
	_ = m.Dao.DeleteSmsCode(rep.PhoneNumber)
	//检查用户是否存在
	exists, err := m.Dao.CheckUserByPhoneNumber(rep.PhoneNumber)
	fmt.Printf("调试: phone=%s, exists=%v, err=%v\n", rep.PhoneNumber, exists, err) // 添加这行
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("用户不存在")
	}
	//获取用户信息
	User, err := m.Dao.LoginUserByPhoneNumber(rep.PhoneNumber)
	if err != nil {
		return nil, err
	}
	//更新登录时间
	now := time.Now()
	User.LastLogin = &now
	//生成token
	token, err := utils.GenerateToken(User.ID, User.UserName)
	if err != nil {
		return nil, err
	}
	//返回更新用户信息
	return &http_models.LoginAndRegister{
		UserName: User.UserName,
		ID:       User.ID,
		Token:    token,
		Avatar:   User.Avatar,
		Expire:   time.Now().Add(time.Hour * 24).Unix(),
	}, nil
}
