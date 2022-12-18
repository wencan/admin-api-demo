package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

// C 配置。
var C config

type config struct {
	// HTTPAddr http服务地址。
	HTTPAddr string `mapstructure:"http_addr"`

	// MySQLDataSourceName MySQL服务的data source name。
	MySQLDataSourceName string `mapstructure:"mysql_datasourcename"`

	// RedisURL redis服务url。
	RedisURL string `mapstructure:"redis_url"`
}

// LoadConfig 加载配置文件。
func LoadConfig(configPath string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed in read config [%s], error: [%w]", configPath, err)
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		return fmt.Errorf("failed in unmarhshal config file [%s], error: [%w]", configPath, err)
	}
	return nil
}
