package provider

type Redis struct {
}

type RedisMethod interface {
}

var Cache = newRedis()

func newRedis() *Redis {
	return &Redis{}
}
