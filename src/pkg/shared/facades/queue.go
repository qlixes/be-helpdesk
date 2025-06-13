package facades

type Queue struct {
}

type QueueManager interface {
}

func NewQueue() *Queue {
	return &Queue{}
}
