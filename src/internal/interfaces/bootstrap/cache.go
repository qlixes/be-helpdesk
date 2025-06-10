package bootstrap

type Cache struct {
}

type CacheManager interface {
}

func NewCache() *Cache {
	return &Cache{}
}
