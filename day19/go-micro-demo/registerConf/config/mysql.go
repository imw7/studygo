package config

import "go-micro.dev/v4/config"

// MySQLConfig 创建结构体
type MySQLConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

// GetMySQLFromConsul 获取mysql的配置
func GetMySQLFromConsul(config config.Config, path ...string) (*MySQLConfig, error) {
	mysqlConfig := &MySQLConfig{}
	// 获取配置
	if err := config.Get(path...).Scan(mysqlConfig); err != nil {
		return nil, err
	}
	return mysqlConfig, nil
}
