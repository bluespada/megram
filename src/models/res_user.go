/**
Copyright (C) 2022 Bluespada, MeGram

This file is under MIT software license, see the accompanying
file COPYING or http://www.opensource.org/licenses/mit-license
*/

package models

import (
	"gorm.io/gorm"
)

type ResUsers struct {
	*gorm.Model
	Username    string `gorm:"unique;not null"`
	PhoneNumber string `gorm:"unique;not null"`
	Name        string
	Picture     string
	Status      string
}
