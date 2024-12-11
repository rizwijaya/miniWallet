package controllers

import "github.com/gofiber/fiber/v2"

func (wc *WalletController) InitMyAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON("")
}
