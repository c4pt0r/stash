package datasource

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

type localDataSource struct {
	m  map[string]Item
	fp *os.File
}

var _ DataSource = &localDataSource{}

func newLocalDataSource() *localDataSource {
	return &localDataSource{
		m: make(map[string]Item),
	}
}

func (l *localDataSource) Get(key []byte) (Item, error) {
	v, ok := l.m[string(key)]
	if !ok {
		return Item{}, ErrNotExist
	}
	return v, nil
}

func (l *localDataSource) List(keyPrfix []byte, opts ListOpt) ([]Item, error) {
	var ret []Item
	for k, v := range l.m {
		if bytes.HasPrefix([]byte(k), keyPrfix) {
			ret = append(ret, v)
		}
	}
	return ret, nil
}

func (l *localDataSource) Put(key []byte, value []byte, meta ItemMeta) error {
	l.m[string(key)] = Item{
		Key:   append([]byte{}, key...),
		Value: append([]byte{}, value...),
		Meta:  meta,
	}
	return nil
}

func (l *localDataSource) Remove(key []byte) error {
	delete(l.m, string(key))
	return nil
}

func (l *localDataSource) Init(path string) error {
	fp, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	l.fp = fp
	body, err := ioutil.ReadAll(l.fp)
	if err != nil {
		fp.Close()
		return err
	}
	if len(body) > 0 {
		err = json.Unmarshal(body, &l.m)
		if err != nil {
			fp.Close()
			return err
		}
	}
	return nil
}

func (l *localDataSource) Sync() error {
	body, err := json.Marshal(l.m)
	if err != nil {
		return err
	}
	l.fp.Truncate(int64(len(body)))
	l.fp.WriteAt(body, 0)
	return l.fp.Sync()
}
