package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Label   string
	Detail  string
	City    string
	OwnerID uint
}
