package facades

type Mail struct {
}

type MailManager interface {
}

func NewMail() *Mail {
	return &Mail{}
}
