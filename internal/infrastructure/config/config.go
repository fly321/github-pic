package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config 配置结构体
type Config struct {
	GitHub struct {
		Token  string `yaml:"token"`
		Cdnurl string `yaml:"cdnurl"`
	}
	Server struct {
		Port string `yaml:"port"`
	}
}

var globalConfig Config

// InitConfig 初始化配置
func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 设置默认值
	viper.SetDefault("server.port", ":8080")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// 如果配置文件不存在，创建默认配置文件
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := createDefaultConfig(); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// 将配置映射到结构体
	return viper.Unmarshal(&globalConfig)
}

// createDefaultConfig 创建默认配置文件
func createDefaultConfig() error {
	defaultConfig := `github:
  token: "your_github_token"
server:
  port: ":8080"
`
	configPath := filepath.Join(".", "config.yaml")
	return os.WriteFile(configPath, []byte(defaultConfig), 0644)
}

// GetConfig 获取全局配置
func GetConfig() *Config {
	fmt.Printf("config: %+v\n", globalConfig)
	return &globalConfig
}
