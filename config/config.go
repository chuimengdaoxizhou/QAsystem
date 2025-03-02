package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Appconfig struct {
	Name        string `mapstructure:"name"`
	Development string `mapstructure:"development"`
	Port        string `mapstructure:"port"`
}

type Jwtconfig struct {
	secretKey string `mapstructure:"secret_key"`
}

type DatabaseConfig struct {
	Driver    string `mapstructure:"driver"`
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Dbname    string `mapstructure:"dbname"`
	Charset   string `mapstructture:"charset"`
	ParseTime string `mapstructure:"parseTime"`
	Loc       string `mapstructure:"loc"`
}

type Config struct {
	App      Appconfig
	Jwt      Jwtconfig
	Database DatabaseConfig
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigFile("yaml")
	viper.AddConfigPath(path)

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return
	}
	return
}

// GetConfig 获取Config文件
func GetConfig() Config {
	config, _ := LoadConfig("./config")
	return config
}

func (c *Config) Getapp() Appconfig {
	appconfig := c.App
	return appconfig
}

func (c *Config) GetDatabasedsn() string {
	dbConfig := c.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Dbname,
		dbConfig.Charset,
		dbConfig.ParseTime,
		dbConfig.Loc,
	)
	return dsn
}

func (c *Config) GetsecretKey() string {
	return c.Jwt.secretKey
}
