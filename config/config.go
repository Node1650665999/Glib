package config

import (
	"github.com/node1650665999/Glib/common"
	"fmt"
	"github.com/gookit/config"
	"github.com/gookit/config/ini"
	"github.com/gookit/config/json"
	"github.com/gookit/config/yaml"
	"github.com/spf13/viper"
)

//ParseConfig2 使用 github.com/gookit/config 来解析配置文件
func ParseConfig2(filename string) (*config.Config, error) {
	drivers := map[string]config.Driver{
		"ini" : ini.Driver,
		"yml" : yaml.Driver,
		"yaml": yaml.Driver,
		"json": json.Driver,
	}

	config.WithOptions(config.ParseEnv)
	ext := common.Ext(filename)
	dr,isset := drivers[ext]
	if ! isset  {
		return nil, fmt.Errorf("file type is invalid, only support [json,yaml,yml,ini]")
	}
	config.AddDriver(dr)
	err := config.LoadFiles(filename)
	if err != nil {
		return nil, fmt.Errorf("load file is err: %v", err)
	}

	return config.Default(), nil
}

//ParseConfig 使用 github.com/spf13/viper 来解析配置文件
func ParseConfig(filename string) (*viper.Viper, error)  {
	viper.SetConfigFile(filename)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("parse config file err: %v", err)
	}
	return viper.GetViper(),nil
}
