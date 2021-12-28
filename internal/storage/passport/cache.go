package passportstorage

import (
	"time"
)

type Cache interface {
	Set(key string, value interface{}, duration time.Duration)
	Get(key string) (interface{}, error)
	Delete(key string) error
}
