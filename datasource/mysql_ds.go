package datasource

import (
	"time"

	"github.com/c4pt0r/log"
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

func newMySQLDatasource(dsn string) *MySQLDatasource {
	return &MySQLDatasource{
		dsn: dsn,
	}
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
	// make sure table exists
	err = db.AutoMigrate(&Items{})
	if err != nil {
		log.E(err)
		return err
	}
	return nil
}

func (m *MySQLDatasource) Sync() error {
	return nil
}
