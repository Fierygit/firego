/*
 * @Author: Firefly
 * @Date: 2021-03-31 13:29:12
 * @Descripttion:
 * @LastEditTime: 2021-03-31 18:28:10
 */

package opt

import (
	"encoding/json"
	"firego/src/common/kv/client"
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	default_opt_len = 6

	kv_opt_key = "opt"

	default_start_cnt = 1
)

type opt_info struct {
	Key  string
	Cnt  int64
	Info string
}

// 帐号信息，可以是 id name，或者相加， key 为 秘钥
func AddOpt(info string, key string) (bool, string) {
	if info == "" || key == "" {
		return false, "帐号或者密码为空"
	}
	data := opt_info{
		Key: key,
		Cnt: int64(default_start_cnt),
	}
	put(info, data)
	return true, ""
}

func GetOpt(info string) (optCode string) {
	if success, data := get(info); success {
		hopt := NewHOTP([]byte(data.Key), default_opt_len)
		// 更新 cnt
		data.Cnt++
		put(info, data)
		return hopt.At(uint64(data.Cnt))
	}
	return ""
}

// 不公开类
func put(key string, info opt_info) (bool, string) {
	data1, err := json.Marshal(info)
	if err != nil {
		logrus.Error("json.marshal failed, err:", err)
		return false, err.Error()
	}
	leveldb := client.NewConnector().SetSize(1).Connect(client.PRE_OPT, "123456") // _opt 是数据库名
	leveldb.Put(kv_opt_key, key, string(data1))                                   // 保存秘钥
	return true, ""
}

func get(key string) (bool, opt_info) {
	leveldb := client.NewConnector().SetSize(1).Connect(client.PRE_OPT, "123456") // _opt 是数据库名
	dataStr := leveldb.Get(kv_opt_key, key)
	ret := opt_info{}
	if dataStr == "" {
		logrus.Info("no user", key)
		return false, ret
	}
	err := json.Unmarshal([]byte(dataStr), &ret)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return false, ret
	}
	return true, ret
}
