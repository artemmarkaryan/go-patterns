package singleton

import "sync"

var once = sync.Once{}
var db DataBase

type DataBase struct {
	url string
}

func GetDataBase() DataBase {
	once.Do(func() {
		db = DataBase{"db://my-db"}
	})
	return db
}