package bootstrap

type Queue struct {
}

type QueueManager interface {
}

func NewQueue() *Queue {
	return &Queue{}
}
