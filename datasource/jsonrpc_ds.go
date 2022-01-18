package datasource

type JsonRPCDatasource struct {
}

var _ DataSource = &JsonRPCDatasource{}

func (j *JsonRPCDatasource) Get(key []byte) (Item, error) {
	panic("not implemented") // TODO: Implement
}

func (j *JsonRPCDatasource) List(keyPrfix []byte, opts ListOpt) ([]Item, error) {
	panic("not implemented") // TODO: Implement
}

func (j *JsonRPCDatasource) Put(key []byte, value []byte, meta ItemMeta) error {
	panic("not implemented") // TODO: Implement
}

func (j *JsonRPCDatasource) Remove(key []byte) error {
	panic("not implemented") // TODO: Implement
}

func (j *JsonRPCDatasource) Init(dsn string) error {
	panic("not implemented") // TODO: Implement
}

func (j *JsonRPCDatasource) Sync() error {
	panic("not implemented") // TODO: Implement
}
