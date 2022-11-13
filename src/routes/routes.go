/**
Copyright (C) 2022 Bluespada, MeGram

This file is under MIT software license, see the accompanying
file COPYING or http://www.opensource.org/licenses/mit-license
*/

package routes

import (
	"github.com/gofiber/fiber/v2"
)

// function handle websocket routing

// Handle http server routing
func InitializeRouter(app *fiber.App) {

	// serve static for web based app
	app.Static("/static", "./static")

	// initialize api
	initializeApi(app.Group("/api"))

	// handle not found
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrNotFound.Code).SendString(fiber.ErrNotFound.Message)
	})
}
