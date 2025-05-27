package domain

type Attachemt struct {
	ID       uint
	ReplyID  uint
	Replies  Reply
	Filename string
}
