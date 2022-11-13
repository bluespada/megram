/**
Copyright (C) 2022 Bluespada, MeGram

This file is under MIT software license, see the accompanying
file COPYING or http://www.opensource.org/licenses/mit-license
*/

package controller

import (
	"github.com/gofiber/fiber/v2"
)

type Auth struct{}

func (a *Auth) HandleLogin(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(map[string]any{
		"error": false,
	})
}

func (a *Auth) HandleRegister(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(map[string]any{})
}
