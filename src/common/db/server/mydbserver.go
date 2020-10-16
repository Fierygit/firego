/*
 * @Author: Firefly
 * @Date: 2020-10-16 19:20:39
 * @Descripttion:
 * @LastEditTime: 2020-10-16 21:41:53
 */
package server

import (
	"firego/src/common/db/model"
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

func (*RPCMethods) Get(req *model.Req, reply *string) error {
	logrus.Info(req)

	err := GetInstance().GetByKey(req.Pair.Key, reply)
	if err != nil {
		logrus.Warn(err)
	}
	return nil
}

func (*RPCMethods) Put(req *model.Req, reply *string) error {
	logrus.Info(req)
	err := GetInstance().PutByKey(req.Pair.Key, req.Pair.Value)
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
