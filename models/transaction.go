package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID    uint
	AddressID uint
	Total     int
	Items     []TransactionItem
}
