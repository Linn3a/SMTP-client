package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"mail/api"
	"mail/model"
)

// 创建一个模型来表示数据库中的数据
//type User struct {
//	gorm.Model
//	Username string
//	Email    string
//}

func main() {
	// 连接到SQLite数据库
	err := model.InitDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 创建一个Fiber应用程序
	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "*",
		}))

	err = api.RegisterRoutes(app)

	if err != nil {
		log.Fatalf("路由注册失败: %v", err)
	}

	err = app.Listen(":8080")
	if err != nil {
		log.Fatalf("Fiber启动失败: %v", err)
	}
}
