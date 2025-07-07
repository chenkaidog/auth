package config

import (
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/yaml.v3"
)

func Init(filepath string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(content, &globalConfig); err != nil {
		panic(err)
	}

	hlog.Debugf("config debug: %+v", globalConfig)
}

func GetMySQLConf() MySQLConf {
	return globalConfig.MySQL
}

func GetRedisConf() RedisConf {
	return globalConfig.Redis
}

func GetJWTConfig() JWTConf {
	return globalConfig.JWT
}

var globalConfig ServiceConf

type ServiceConf struct {
	MySQL MySQLConf `yaml:"mysql"`
	Redis RedisConf `yaml:"redis"`
	JWT   JWTConf   `yaml:"jwt"`
}

type MySQLConf struct {
	DBName   string `yaml:"db_name"`
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RedisConf struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type JWTConf struct {
	Issuer     string `yaml:"issuer"`
	SecretKey  string `yaml:"secret_key"`
	Expiration int    `yaml:"expiration"`
}
