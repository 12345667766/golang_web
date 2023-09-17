package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var C Cfg

type Cfg struct {
	RedisConfig
	ServerConfig
	LogConfig
}

type LogConfig struct {
	DebugFileName string
	InfoFileName  string
	WarnFileName  string
	MaxSize       int
	MaxAge        int
	MaxBackups    int
}

type RedisConfig struct {
	Address  string
	Password string
	Db       int
}

type ServerConfig struct {
	ServerName string
	Address    string
}

func InitConfig() {
	vip := viper.New()
	cfgPath, _ := os.Getwd()
	vip.SetConfigName("config")
	vip.SetConfigType("yaml")
	vip.AddConfigPath(cfgPath + "\\..\\project-common\\config")
	vip.AddConfigPath("/etc/project_user")
	err := vip.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件出错: ", err)
	}
	readRedisConfig(vip)
	readServer(vip)
	readLogsConfig(vip)
}

func readLogsConfig(vip *viper.Viper) {
	C.LogConfig.DebugFileName = vip.GetString("logs.DebugFileName")
	C.LogConfig.InfoFileName = vip.GetString("logs.InfoFileName")
	C.LogConfig.WarnFileName = vip.GetString("logs.WarnFileName")
	C.LogConfig.MaxSize = vip.GetInt("MaxSize")
	C.LogConfig.MaxSize = vip.GetInt("MaxAge")
	C.LogConfig.MaxSize = vip.GetInt("MaxBackups")
}

func readRedisConfig(vip *viper.Viper) {
	C.RedisConfig.Address = vip.GetString("redis.address")
	C.RedisConfig.Password = vip.GetString("redis.password")
	C.RedisConfig.Db = vip.GetInt("db")
}

func readServer(vip *viper.Viper) {
	C.ServerConfig.ServerName = vip.GetString("server.serverName")
	C.ServerConfig.Address = vip.GetString("server.address")
}
