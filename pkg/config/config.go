package config

import (
	log "github.com/cihub/seelog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

// 全局配置变量
var Conf = new(config)

// 设置读取配置信息
func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Errorf("读取应用目录失败:%s", err)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/")
	// 读取配置信息
	err = viper.ReadInConfig()

	// 热更新配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 将读取的配置信息保存至全局变量Conf
		if err := viper.Unmarshal(Conf); err != nil {
			log.Errorf("初始化配置文件失败:%s", err)
		}
		Conf.System.RSAPublicBytes = RSAReadKeyFromFile(Conf.System.RSAPublicKey)
		Conf.System.RSAPrivateBytes = RSAReadKeyFromFile(Conf.System.RSAPrivateKey)
	})

	if err != nil {
		log.Errorf("读取配置文件失败:%s", err)
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		log.Errorf("初始化配置文件失败:%s", err)
	}

	// 读取rsa key
	Conf.System.RSAPublicBytes = RSAReadKeyFromFile(Conf.System.RSAPublicKey)
	Conf.System.RSAPrivateBytes = RSAReadKeyFromFile(Conf.System.RSAPrivateKey)
}

type config struct {
	System *SystemConfig `mapstructure:"system" json:"system"`
	Logs   *LogsConfig   `mapstructure:"logs" json:"logs"`
	Mysql  *MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	//Casbin    *CasbinConfig    `mapstructure:"casbin" json:"casbin"`
	Jwt       *JwtConfig       `mapstructure:"jwt" json:"jwt"`
	RateLimit *RateLimitConfig `mapstructure:"rate-limit" json:"rateLimit"`
}

type SystemConfig struct {
	Mode            string `mapstructure:"mode" json:"mode"`
	UrlPathPrefix   string `mapstructure:"url-path-prefix" json:"urlPathPrefix"`
	Port            int    `mapstructure:"port" json:"port"`
	InitData        bool   `mapstructure:"init-data" json:"initData"`
	RSAPublicKey    string `mapstructure:"rsa-public-key" json:"rsaPublicKey"`
	RSAPrivateKey   string `mapstructure:"rsa-private-key" json:"rsaPrivateKey"`
	RSAPublicBytes  []byte `mapstructure:"-" json:"-"`
	RSAPrivateBytes []byte `mapstructure:"-" json:"-"`
}

type MysqlConfig struct {
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	Database    string `mapstructure:"database" json:"database"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Query       string `mapstructure:"query" json:"query"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode"`
	TablePrefix string `mapstructure:"table-prefix" json:"tablePrefix"`
	Charset     string `mapstructure:"charset" json:"charset"`
	Collation   string `mapstructure:"collation" json:"collation"`
}

type LogsConfig struct {
	LogConfigStr string `mapstructure:"log-config-str" json:"logconfigstr"`
}

type RateLimitConfig struct {
	FillInterval int64 `mapstructure:"fill-interval" json:"fillInterval"`
	Capacity     int64 `mapstructure:"capacity" json:"capacity"`
}

type JwtConfig struct {
	Realm      string `mapstructure:"realm" json:"realm"`
	Key        string `mapstructure:"key" json:"key"`
	Timeout    int    `mapstructure:"timeout" json:"timeout"`
	MaxRefresh int    `mapstructure:"max-refresh" json:"maxRefresh"`
}

// 从文件中读取RSA key
func RSAReadKeyFromFile(filename string) []byte {
	f, err := os.Open(filename)
	var b []byte

	if err != nil {
		return b
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b = make([]byte, fileInfo.Size())
	_, err = f.Read(b)
	if err != nil {
		return b
	}
	return b
}
