package storage

import (
	"io"
)

type Storage interface {
	Put(key string, r io.Reader, dataLength int64) error
	PutFile(key string, localFile string) error
	Get(key string) (io.ReadCloser, error)
	Rename(srcKey string, destKey string) error
	Copy(srcKey string, destKey string) error
	Exists(key string) (bool, error)
	Size(key string) (int64, error)
	Delete(key string) error
	Url(key string) string
	IsLocal() bool
}
