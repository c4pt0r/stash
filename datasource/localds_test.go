package datasource

import (
	"fmt"
	"testing"
	"time"
)

func TestLocalDsInit(t *testing.T) {
	ds, err := Init("local://local")
	if err != nil {
		t.Error(err)
	}

	ds.Put([]byte("hello"), []byte("world"), ItemMeta{
		Tp:        TypeBytes,
		Len:       len([]byte("world")),
		CreateAt:  time.Now(),
		Tags:      nil,
		Namespace: DefaultNS,
	})

	fmt.Println(ds.Get([]byte("hello")))

	ds.Put([]byte("a/1"), []byte("1"), ItemMeta{
		Tp:        TypeBytes,
		Len:       len([]byte("1")),
		CreateAt:  time.Now(),
		Tags:      nil,
		Namespace: DefaultNS,
	})

	ds.Put([]byte("a/2"), []byte("2"), ItemMeta{
		Tp:        TypeBytes,
		Len:       len([]byte("2")),
		CreateAt:  time.Now(),
		Tags:      nil,
		Namespace: DefaultNS,
	})

	ds.Put([]byte("a/3"), []byte("3"), ItemMeta{
		Tp:        TypeBytes,
		Len:       len([]byte("3")),
		CreateAt:  time.Now(),
		Tags:      nil,
		Namespace: DefaultNS,
	})

	out, _ := ds.List([]byte("a/"), nil)
	fmt.Println(out)

}
