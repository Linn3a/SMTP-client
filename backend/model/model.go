package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Username string
	Address  string
}

type Mail struct {
	gorm.Model
	Subject string
	Content string
	To      []*User `gorm:"many2many:mail_to"`
}

type Craft struct { // 将 craft 的名称首字母改成大写
	gorm.Model
	Subject string
	Content string
	To      []*User `gorm:"many2many:craft_to"`
}

func InitDB() error {
	db, err := gorm.Open(sqlite.Open("email.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
		return err
	}
	DB = db

	// 自动迁移模型
	err = DB.AutoMigrate(&User{}, &Mail{}, &Craft{})
	if err != nil {
		return err
	}
	return nil
}
