package valueobject

import "gorm.io/gorm"

type FaqAttachment struct {
	gorm.Model
	ID       uint
	FaqID    uint
	Faqs     Faq
	Filename string
}
