package domain

type Reply struct {
	ID       uint
	TicketID uint
	Tickets  Ticket
	UserID   uint
	Users    User
	Body     string
}
