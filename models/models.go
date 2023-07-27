package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Sender  string `gorm:"column:sender"`
	Content string `gorm:"column:content"`
	Body    string `gorm:"column:body"`
}
