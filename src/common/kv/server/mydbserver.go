/*
 * @Author: Firefly
 * @Date: 2020-10-16 19:20:39
 * @Descripttion:
 * @LastEditTime: 2021-04-23 16:57:26
 */
package server

import (
	"firego/src/common/kv/model"
	"net"
	"net/rpc"
	"os"

	"github.com/sirupsen/logrus"
)

//Run start
func Run() {

	InitDB()
	initServer()

}

//RPCMethods proxy
type RPCMethods int

//Get g
func (*RPCMethods) Get(req *model.Req, reply *string) error {
	logrus.Info("Get ", req)
	err := GetInstance().GetByKey(req.Pair.Key, reply)
	if err != nil {
		logrus.Warn(err)
	}
	return nil
}

//BatchGet g
func (*RPCMethods) BatchGet(req *model.Req, reply *[]string) error {
	logrus.Info("BatchGet ", req)
	GetInstance().BatchGetByPrefix(req.Pair.Key, reply)
	return nil
}

//Put p
func (*RPCMethods) Put(req *model.Req, reply *string) error {
	logrus.Info("Put ", req)
	err := GetInstance().PutByKey(req.Pair.Key, req.Pair.Value)
	if err != nil {
		logrus.Warn(err)
	}
	return nil
}

//Has p
func (*RPCMethods) Has(req *model.Req, reply *bool) error {
	logrus.Info("Has ", req)
	err := GetInstance().Has(req.Pair.Key, reply)
	if err != nil {
		logrus.Warn(err)
	}
	return nil
}

//Delete p
func (*RPCMethods) Delete(req *model.Req, reply *string) error {
	logrus.Info("delete  ", req)
	err := GetInstance().Delete(req.Pair.Key)

	if err != nil {
		logrus.Warn(err)
	}
	return nil
}

func initServer() {
	methods := new(RPCMethods)
	logrus.Info(methods)
	err := rpc.Register(methods)
	checkError(err)

	addy, err := net.ResolveTCPAddr("tcp", ":3307")
	checkError(err)

	// 先监听, 再连接
	listener, err := net.ListenTCP("tcp", addy)
	checkError(err)

	logrus.Infof("rpc server init over")
	for {
		conn, _ := listener.Accept()
		logrus.Infof("receive from %s", conn.RemoteAddr())
		go rpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		logrus.Info("Fatal error ", err.Error())
		os.Exit(1)
	}
}
