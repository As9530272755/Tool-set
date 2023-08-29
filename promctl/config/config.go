package config

import "github.com/spf13/viper"

// 定义 log 结构体
type Log struct {
	FileName string `mapstructure:"filename"`
	Max_age  int    `mapstructure:"max_age"`
	Max_size int    `mapstructure:"max_size"`
	Compress bool   `mapstructure:"compress"`
	Level    string `mapstructure:"level"`
}

// 定义 web 结构体
type Prometheus struct {
	Addr string `mapstructuer:"addr"`
	Port string `mapstructuer:"port"`
}

// 定义解析的配置文件对象
type Options struct {
	Prometheus Prometheus `mapstructuer:"prometheus"`
	Log        Log        `mapstructuer:"log"`
}

// 解析配置文件并返回解析后的 options 结构体
func ParseConfig() (*Options, error) {
	// 定义一个新对象
	conf := viper.New()

	// 解析传入的配置文件
	conf.SetConfigFile("./etc/config.yaml")

	// 解析配置文件
	if err := conf.ReadInConfig(); err != nil {
		return nil, err
	}

	// 将配置文件内容读取到结构体中
	option := &Options{}
	err := conf.Unmarshal(&option)
	if err != nil {
		return nil, err
	}
	return option, nil
}
