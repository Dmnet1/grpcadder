package cache

type Storage interface {
	Set(key, value string) (err error)
	Get(key string) (valuer string, err error)
	Delete(key string) (err error)
}
