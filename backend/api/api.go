package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"mail/mailutils"
	"mail/model"
)

func RegisterRoutes(app *fiber.App) error {
	type userWhenCreate struct {
		Username string `json:"username"`
		Address  string `json:"address"`
	}
	type CreateMailRequest struct {
		Subject string           `json:"subject"`
		Content string           `json:"content"`
		To      []userWhenCreate `json:"to"`
	}

	// 登录
	type loginBody struct {
		User     string `json:"username"`
		Password string `json:"password"`
	}
	app.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()
		var login loginBody
		err := json.Unmarshal(body, &login)
		fmt.Printf("login: %+v\n", login)
		fmt.Println(login.User == mailutils.User, login.Password == mailutils.Password)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status": "parse error",
			})
		}
		if login.User == mailutils.User && login.Password == mailutils.Password {
			return c.Status(200).JSON(fiber.Map{
				"status": "ok",
			})
		} else {
			return c.Status(401).JSON(fiber.Map{
				"status": "error",
			})
		}
	})
	// 发送邮件
	app.Post("/send", func(c *fiber.Ctx) error {
		//	获取body
		body := c.Body()
		//	解析body
		var newmail CreateMailRequest
		err := json.Unmarshal(body, &newmail)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err,
			})
		}
		fmt.Printf("newmail: %+v\n", newmail)
		Users := make([]*model.User, len(newmail.To))
		for index, user := range newmail.To {
			var savedUser model.User
			fmt.Printf("user: %+v\n", user)
			err = model.DB.First(&savedUser, "address = ?", user.Address).Error
			fmt.Printf("savedUser: %v", savedUser)
			fmt.Print(savedUser == model.User{})
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Printf("not found\n")
				err = model.DB.Create(&model.User{
					Address:  user.Address,
					Username: user.Username,
				}).Error
				if err != nil {
					return c.Status(500).JSON(fiber.Map{
						"error": err,
					})
				}

				err = model.DB.First(&savedUser, "address = ?", user.Address).Error
				if err != nil {
					return c.Status(500).JSON(fiber.Map{
						"error": err,
					})
				}
				fmt.Printf("saving user: %v\n", savedUser)
			}
			fmt.Printf("saving user: %v\n", savedUser)
			Users[index] = &savedUser
		}

		mailutils.SendMail(
			Users,
			newmail.Subject,
			newmail.Content,
		)
		savedMail := model.Mail{
			Subject: newmail.Subject,
			Content: newmail.Content,
			To:      Users,
		}

		model.DB.Create(&savedMail)

		return c.Status(200).JSON(fiber.Map{
			"message": "邮件发送成功",
			"mail":    savedMail,
		})
	})

	// 获取所有邮件
	app.Get("/mails", func(c *fiber.Ctx) error {
		var mails []*model.Mail
		err := model.DB.Preload("To").Find(&mails).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"mails": mails,
		})
	})

	// 获取通讯录
	app.Get("/users", func(c *fiber.Ctx) error {
		var users []*model.User
		err := model.DB.Find(&users).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"users": users,
		})
	})

	// 获取所有草稿
	app.Get("/crafts", func(c *fiber.Ctx) error {
		var crafts []*model.Craft
		err := model.DB.Preload("To").Find(&crafts).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"crafts": crafts,
		})
	})

	// 获取单个草稿
	app.Get("/crafts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var craft model.Craft
		err := model.DB.Preload("To").First(&craft, id).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"craft": craft,
		})
	})

	// 保存草稿
	app.Post("/crafts", func(c *fiber.Ctx) error {
		body := c.Body()
		//	解析body
		var newCraft CreateMailRequest
		err := json.Unmarshal(body, &newCraft)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err,
			})
		}
		Users := make([]*model.User, len(newCraft.To))
		for index, user := range newCraft.To {
			var savedUser model.User
			fmt.Printf("user: %+v\n", user)
			err = model.DB.First(&savedUser, "address = ?", user.Address).Error
			fmt.Printf("savedUser: %v", savedUser)
			fmt.Print(savedUser == model.User{})
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Printf("not found\n")
				err = model.DB.Create(&model.User{
					Address:  user.Address,
					Username: user.Username,
				}).Error
				if err != nil {
					return c.Status(500).JSON(fiber.Map{
						"error": err,
					})
				}

				err = model.DB.First(&savedUser, "address = ?", user.Address).Error
				if err != nil {
					return c.Status(500).JSON(fiber.Map{
						"error": err,
					})
				}
				fmt.Printf("saving user: %v\n", savedUser)
			}
			fmt.Printf("saving user: %v\n", savedUser)
			Users[index] = &savedUser
		}

		savedCraft := model.Craft{
			Subject: newCraft.Subject,
			Content: newCraft.Content,
			To:      Users,
		}

		model.DB.Create(&savedCraft)

		return c.Status(200).JSON(fiber.Map{
			"message": "草稿保存成功",
			"mail":    savedCraft,
		})

	})

	// 更新草稿
	app.Post("/crafts/:id", func(c *fiber.Ctx) error {
		body := c.Body()
		id := c.Params("id")
		//	解析body
		var newCraft CreateMailRequest
		err := json.Unmarshal(body, &newCraft)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err,
			})
		}
		Users := make([]*model.User, len(newCraft.To))
		for index, user := range newCraft.To {
			var savedUser model.User
			fmt.Printf("user: %+v\n", user)
			err = model.DB.First(&savedUser, "address = ?", user.Address).Error
			fmt.Printf("savedUser: %v", savedUser)
			fmt.Print(savedUser == model.User{})
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Printf("not found\n")
				err = model.DB.Create(&model.User{
					Address:  user.Address,
					Username: user.Username,
				}).Error
				if err != nil {
					return c.Status(500).JSON(fiber.Map{
						"error": err,
					})
				}

				err = model.DB.First(&savedUser, "address = ?", user.Address).Error
				if err != nil {
					return c.Status(500).JSON(fiber.Map{
						"error": err,
					})
				}
				fmt.Printf("saving user: %v\n", savedUser)
			}
			fmt.Printf("saving user: %v\n", savedUser)
			Users[index] = &savedUser
		}

		savedCraft := model.Craft{
			Subject: newCraft.Subject,
			Content: newCraft.Content,
			To:      Users,
		}

		model.DB.Model(&model.Craft{}).Where("id = ?", id).Updates(&savedCraft)

		return c.Status(200).JSON(fiber.Map{
			"message": "草稿更新成功",
			"mail":    savedCraft,
		})

	})

	return nil
}
