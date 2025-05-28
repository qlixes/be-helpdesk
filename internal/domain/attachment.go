package domain

import "gorm.io/gorm"

type Attachemt struct {
	gorm.Model
	ID       uint
	ReplyID  uint
	Replies  Reply
	Filename string
}
