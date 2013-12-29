package common_layer

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/cache"
	"github.com/syndtr/goleveldb/leveldb"
)

type myleveldb struct {
	*leveldb.DB
}

func (*myleveldb) ClearAll() error {
	return errors.New("ClearAll method not implement")
}

func (*myleveldb) Decr(str string) error {
	return errors.New("Decr method not implement")
}

func (*myleveldb) Delete(str string) error {
	return errors.New("Delete method not implement")
}

func (*myleveldb) Get(str string) interface{} {
	return errors.New("Get method not implement")
}

func (*myleveldb) Incr(str string) error {
	return errors.New("Incr method not implement")
}

func (*myleveldb) IsExist(str string) bool {
	return false
}

func (*myleveldb) Put(str string, inter interface{}, key int64) error {
	return errors.New("Put method not implement")
}

func (*myleveldb) StartAndGC(str string) error {
	return errors.New("StartAndGC method not implement")
}

func init() {
	fmt.Println("init cache memory")
	cache.Register("leveldb", Using_level())
}

func Using_cache() {
	bm, err := cache.NewCache("memory", `{"interval":60},{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120},{"conn":"127.0.0.1:11211"}`)
	fmt.Println(err)
	fmt.Println(bm.Get("key"))
}
func Using_level() *myleveldb {

	var mydb myleveldb
	var err error
	mydb.DB, err = leveldb.OpenFile("E:/db", nil)
	fmt.Println(err)
	defer mydb.DB.Close()
	data, err := mydb.DB.Get([]byte("key"), nil)
	fmt.Println(data)
	err = mydb.DB.Put([]byte("key"), []byte("value"), nil)

	err = mydb.DB.Delete([]byte("key"), nil)

	iter := mydb.DB.NewIterator(nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()

		value := iter.Value()
		fmt.Println(key)
		fmt.Println(value)
	}
	//iter.Release()
	//err = iter.Error()
	//batch := new(leveldb.Batch)
	//batch.Put([]byte("foo"), []byte("value"))
	//batch.Put([]byte("bar"), []byte("another value"))
	//batch.Delete([]byte("baz"))
	//err = db.Write(batch, nil)

	//o := &opt.Options{
	//	Filter: filter.NewBloomFilter(10),
	//}
	//db, err := leveldb.OpenFile("E:/db", o)

	return &mydb
}
