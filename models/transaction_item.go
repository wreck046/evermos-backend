package models

import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model
	TransactionID uint
	ProductID     uint
	ProductName   string
	Price         int
	Qty           int
	Subtotal      int
}
