package domain

type FaqAttachment struct {
	ID       uint
	FaqID    uint
	Faqs     Faq
	Filename string
}
