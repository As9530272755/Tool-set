package config

import (
	"github.com/spf13/viper"
)

type KubeConfig struct {
	ConfigPath string `mapstructure:"path"`
}

type Log struct {
	FileName string `mapstructure:"filename"`
	Max_age  int    `mapstructure:"max_age"`
	Max_size int    `mapstructure:"max_size"`
	Compress bool   `mapstructure:"compress"`
	Level    string `mapstructure:"level"`
}

type Optins struct {
	KubeConfig KubeConfig `mapstructure:"kubeconfig"`
	Log        Log        `mapstructure:"log"`
}

// 指定配置文件路径,并返回 optins 结构体
func ParseConfig() (*Optins, error) {
	conf := viper.New()
	conf.SetConfigFile("./etc/saConfig.yaml")

	if err := conf.ReadInConfig(); err != nil {
		return nil, err
	}

	optins := &Optins{}
	if err := conf.Unmarshal(&optins); err != nil {
		return nil, err
	}
	return optins, nil
}
