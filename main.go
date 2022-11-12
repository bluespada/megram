/**
Copyright (C) 2022 Bluespada, MeGram

This file is under MIT software license, see the accompanying
file COPYING or http://www.opensource.org/licenses/mit-license
*/

package main

import (
	"log"
	"os"
	"sync"

	ServerUtils "github.com/bluespada/megram/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// load env
	godotenv.Load()

	var wg = new(sync.WaitGroup)
	app := fiber.New(fiber.Config{
		AppName: "MeGram Server",
	})

	wg.Add(2)

	go func() {
		// initialize database
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  os.Getenv("DB_DSN"),
			PreferSimpleProtocol: true,
		}))
		if err != nil {
			log.Fatalln(err)
		}
		// initialize auto migrate for models

		ServerUtils.AppDatabase = db
		wg.Done()
	}()

	go func() {
		// initialize server
		log.Fatalln(app.Listen(":" + os.Getenv("APP_PORT")))
		wg.Done()
	}()

	// wait
	wg.Wait()
}
