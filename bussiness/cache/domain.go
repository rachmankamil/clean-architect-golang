package cache

type Repository interface {
	Get(key string) string
	Set(key, value string) string
}
