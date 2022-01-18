package datasource

import (
	"errors"
	"strings"
	"time"
)

type ItemType int

const (
	TypeBytes ItemType = iota + 1
	TypeJSON
)

const (
	DefaultNS = "default"
)

var (
	ErrUnknownProtocol error = errors.New("unknown protocol")
	ErrNotExist        error = errors.New("no such item")
)

type ItemMeta struct {
	Tp        ItemType
	Len       int
	CreateAt  time.Time
	Tags      []string
	Namespace string
}

type Item struct {
	Key   []byte
	Value []byte
	Meta  ItemMeta
}

type ListOpt interface{}

type DataSource interface {
	Get(key []byte) (Item, error)
	List(keyPrfix []byte, opts ListOpt) ([]Item, error)
	Put(key, value []byte, meta ItemMeta) error
	Remove(key []byte) error
	Init(dsn string) error
	Sync() error
}

func Init(dsn string) (DataSource, error) {
	parts := strings.SplitN(dsn, "://", 2)
	if len(parts) != 2 {
		return nil, ErrUnknownProtocol
	}
	switch strings.ToLower(parts[0]) {
	case "local":
		localDs := newLocalDataSource()
		err := localDs.Init(parts[1])
		if err != nil {
			return nil, err
		}
		return localDs, nil
	}
	return nil, ErrUnknownProtocol
}
