package repository

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	// 取設定檔
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 連線資料庫資料
	user := viper.GetString("mysql.user")
	pass := viper.GetString("mysql.pass")
	ip := viper.GetString("mysql.ip")
	dbname := viper.GetString("mysql.db")

	// 執行連線
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, ip, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
