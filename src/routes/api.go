/**
Copyright (C) 2022 Bluespada, MeGram

This file is under MIT software license, see the accompanying
file COPYING or http://www.opensource.org/licenses/mit-license
*/

package routes

import (
	ServerController "github.com/bluespada/megram/src/controller"
	"github.com/gofiber/fiber/v2"
)

func initializeApi(api fiber.Router) {
	auth := ServerController.Auth{}
	// handle login
	apiAuth := api.Group("/auth")
	apiAuth.Post("/login", auth.HandleLogin)
}
