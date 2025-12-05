package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server string `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	JWT    JWT    `yaml:"jwt"`
}

type Mysql struct {
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
}
type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}
type JWT struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

// LoadConfig加载配置文件
func LoadConfig(path string) (*Config, error) {
	var config Config
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &config) //将字节切片data解析为config结构体的指针
	if err != nil {
		return nil, err
	}
	return &config, nil
}
