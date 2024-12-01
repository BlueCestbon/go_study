package main

import (
	"github.com/rosedblabs/rosedb/v2"
	"log"
)

func main() {
	// 指定选项
	options := rosedb.DefaultOptions
	options.DirPath = "E:\\Go_Project\\study\\rosedb_x\\rosedb_basic"

	// 打开数据库
	db, err := rosedb.Open(options)
	if err != nil {
		panic(err)
	}
	defer func() {
		log.Println("关闭rosedb")
		_ = db.Close()
	}()

	// 设置键值对
	err = db.Put([]byte("name"), []byte("rosedb"))
	if err != nil {
		panic(err)
	}

	// 获取键值对
	val, err := db.Get([]byte("name"))
	if err != nil {
		panic(err)
	}
	println(string(val))

	//// 删除键值对
	//err = db.Delete([]byte("name"))
	//if err != nil {
	//	panic(err)
	//}

	// 再次获取键值对
	//val, err = db.Get([]byte("name"))
	//if err != nil {
	//	panic(err)
	//}
	//println(string(val))
}
