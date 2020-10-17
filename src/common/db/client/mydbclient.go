package client

import (
	"firego/src/common/db/model"
	"math/rand"
	"net/rpc"

	"github.com/sirupsen/logrus"
)

//Leveldb db
type Leveldb struct {
	con      *Connector
	name     string //`db name and the username`
	password string //`password to the db`
}

//Connector engin
type Connector struct {
	conSize int //连接大小
	cons    []*rpc.Client
	//keepalive
}

//-------------------------------connector-----------------

//NewConnector constructer
func NewConnector() *Connector {
	return &Connector{
		conSize: 2,
	}
}

//SetSize poolsize
func (con *Connector) SetSize(size int) *Connector {
	con.conSize = size
	return con
}

//GetSize getsize
func (con *Connector) GetSize() int {
	return con.conSize
}

//Connect connector
func (con *Connector) Connect(name string, password string) (leveldb Leveldb) {
	con.cons = make([]*rpc.Client, con.conSize)
	//rpc的与服务端建立网络连接
	for i := range con.cons {
		i := i
		cli, err := rpc.Dial("tcp", "127.0.0.1:3307")
		if err != nil {
			panic(err)
		}
		con.cons[i] = cli
	}
	logrus.Info("init ", con.conSize, " connection")
	for i, j := range con.cons {
		logrus.Info(i, j)
	}
	return Leveldb{password: password, name: name, con: con}
}

//GetCon g
func (con *Connector) GetCon() *rpc.Client {
	return con.cons[rand.Intn(con.GetSize())]
}

// ----------------------Leveldb--------------------------------

//Get g
func (db *Leveldb) Get(key interface{}) interface{} {
	// 1、类型检查， 只有string！！！
	// todo 自定义序列化
	keyTmp := key.(string)
	valueTmp := string("")
	return *(db.call(&keyTmp, &valueTmp, "get")).(*string)
}

//Put put
func (db *Leveldb) Put(key interface{}, value interface{}) {
	//todo 自定义序列化
	keyTmp := key.(string)
	valueTmp := value.(string)
	db.call(&keyTmp, &valueTmp, "put")
}

//Has put
func (db *Leveldb) Has(key interface{}) bool {
	//todo 自定义序列化
	keyTmp := key.(string)
	valueTmp := string("")
	return db.call(&keyTmp, &valueTmp, "has").(bool)
}

//Delete put
func (db *Leveldb) Delete(key interface{}, value interface{}) {
	//todo 自定义序列化
	keyTmp := key.(string)
	valueTmp := value.(string)
	db.call(&keyTmp, &valueTmp, "del")
}

//Put0 put
func (db *Leveldb) call(key *string, value *string, opt string) interface{} {
	rep := ""
	pairTmp := model.NewPair(db.name+*key, *value)
	var rst error
	switch opt {
	case "del":
		rst = db.con.GetCon().Call("RPCMethods.Delete", model.NewReq(&db.name, &db.password, pairTmp), &rep)
	case "get":
		rst = db.con.GetCon().Call("RPCMethods.Get", model.NewReq(&db.name, &db.password, pairTmp), &rep)
	case "put":
		rst = db.con.GetCon().Call("RPCMethods.Put", model.NewReq(&db.name, &db.password, pairTmp), &rep)
	case "has":
		var reply bool
		rst = db.con.GetCon().Call("RPCMethods.Has", model.NewReq(&db.name, &db.password, pairTmp), &reply)
		if rst != nil {
			//todo 连接错误处理
			panic("con error")
		}
		return reply
	}
	if rst != nil {
		//todo 连接错误处理
		panic("con error")
	}
	return &rep
}

//Get1 put
func (db *Leveldb) Get1(key1 interface{}, key2 interface{}) {
	keyTmp := key1.(string) + key2.(string)
	db.Get(keyTmp)
}

//Put1 put
func (db *Leveldb) Put1(key1 interface{}, key2 interface{}, value interface{}) {
	keyTmp := key1.(string) + key2.(string)
	db.Put(keyTmp, value)
}
