/*
 * @Author: Firefly
 * @Date: 2020-10-16 16:49:57
 * @Descripttion:
 * @LastEditTime: 2020-10-16 20:40:57
 */
package model

type Pair struct {
	Key   string
	Value string
}

type Req struct {
	Name      string
	Passworld string
	Pair      Pair
}

//NewReq req
func NewReq(name *string, passworld *string, pair Pair) Req {
	return Req{
		Name:      *name,
		Passworld: *passworld,
		Pair:      pair,
	}
}

//NewPair 参数上创建了对象
func NewPair(key string, value string) Pair {
	return Pair{
		Key:   key,
		Value: value,
	}
}
