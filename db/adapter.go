package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gobase/config"
)

var mysqlConnector *MySQLConnector

type DBConnectAdapter interface {
	Connect() (*gorm.DB, error)
}

type MySQLConnector struct {
	config.MySQLConf
	DB *gorm.DB
}

func GetMySQLConnector() *MySQLConnector {
	appConfig := config.GetConfig()
	if mysqlConnector == nil {
		mysqlConnector = &MySQLConnector{
			MySQLConf: config.MySQLConf{
				Host:          appConfig.MySQL.Host,
				Port:          appConfig.MySQL.Port,
				Username:      appConfig.MySQL.Username,
				Password:      appConfig.MySQL.Password,
				DefaultDB:     appConfig.MySQL.DefaultDB,
				MaxConnection: appConfig.MySQL.MaxConnection,
			},
		}
	}
	return mysqlConnector
}

func (conn *MySQLConnector) Connect() (*gorm.DB, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conn.Username, conn.Password, conn.Host, conn.Port, conn.DefaultDB)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	mysqlConnector.DB = db
	return db, nil
}

func GetDB() *gorm.DB {
	return mysqlConnector.DB
}
