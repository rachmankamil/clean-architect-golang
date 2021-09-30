package cache

type Repository interface {
	Get(key string) (string, error)
	Set(key, value string) (string, error)
	Remove(key string) (string, error)
}
