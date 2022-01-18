package datasource

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	Key          string // primary key
	JsonContent  []byte // TODO, json column
	BytesContent []byte
	Namespace    string   // namespace
	Tags         []string // index
	CreateAt     time.Time
}

type MySQLDatasource struct {
	dsn string
	db  *gorm.DB
}

func (m *MySQLDatasource) Get(key []byte) (Item, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MySQLDatasource) List(keyPrfix []byte, opts ListOpt) ([]Item, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MySQLDatasource) Put(key []byte, value []byte, meta ItemMeta) error {
	panic("not implemented") // TODO: Implement
}

func (m *MySQLDatasource) Remove(key []byte) error {
	panic("not implemented") // TODO: Implement
}

func (m *MySQLDatasource) Init(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	m.db = db
	m.dsn = dsn
	return nil
}

func (m *MySQLDatasource) Sync() error {
	return nil
}
