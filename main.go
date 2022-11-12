/**
Copyright (C) 2022 Bluespada, MeGram

This file is under MIT software license, see the accompanying
file COPYING or http://www.opensource.org/licenses/mit-license
*/

package main

import (
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	// load env
	godotenv.Load()

	var wg = new(sync.WaitGroup)

	// wait
	wg.Wait()
}
