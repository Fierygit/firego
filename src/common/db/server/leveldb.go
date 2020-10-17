/*
 * @Author: Firefly
 * @Date: 2020-10-16 15:24:48
 * @Descripttion:
 * @LastEditTime: 2020-10-17 15:51:20
 */
package server

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

//InitDB s
func InitDB() {
	logrus.Info("init databse...")
	// InitDataBase 初始化数据库
	db, err := leveldb.OpenFile("../data", nil)
	if err != nil {
		os.Exit(1)
	}

	instance = &DataBase{db: db}
}

// DataBase 对leveldb的简单封装
type DataBase struct {
	db *leveldb.DB
}

// DB 全局对象
var instance *DataBase

// GetInstance 获取单例对象
func GetInstance() *DataBase {
	return instance
}

// --------------------- DataBase operation -----------------------

// GetByKey get value by key
func (p *DataBase) GetByKey(key string, value *string) error {
	data, err := p.db.Get([]byte(key), nil)
	*value = string(data)
	return err
}

// BatchGetByPrefix 通过前缀获取数据
func (p *DataBase) BatchGetByPrefix(prefix string, values []interface{}) {
	iter := p.db.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	for iter.Next() {
		values = append(values, iter.Value)
	}
	iter.Release()
}

// PutByKey 存放数据
func (p *DataBase) PutByKey(key string, value string) error {
	err := p.db.Put([]byte(key), []byte(value), nil)
	return err
}

// Has d
func (p *DataBase) Has(key string, value *bool) error {
	data, err := p.db.Has([]byte(key), nil)
	*value = data
	return err
}

// Delete d
func (p *DataBase) Delete(key string) error {
	return p.db.Delete([]byte(key), nil)
}
