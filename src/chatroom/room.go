package chatroom

type Client struct {
	IP            string // 用户的ip
	Name          string // 用户的昵称， 一个要一个， 前段可以随机生成
	Index         int    // 已经拉去消息的索引位置
	IsAlive       bool   // 这个用户时候还在
	LastAliveTime int64  // 上一次在线的时间
}

type Message struct {
	Type     int         // 消息类型， 0 为普通消息， 1 为提示消息等等
	Info     interface{} // 存储消息
	SendTime int64       // 存储这条消息的时间
}

type RoomInfo struct {
	CreateInfo CreateRoomArg // 创建房间必要的参数
	CreateTime int64         // 创建的时间
}

type Room struct {
	RoomInfo  RoomInfo
	ClientNum int               //	在线的人数
	Messages  []Message         // 	消息列表， 索引号按时间排序
	Clients   map[string]string //	SessionId -> Client
}

// chatroom 的结构体
// 暂时不支持通过房间名字查询
type ChatRoom struct {
	RoomNum int             // 房间的个数
	Rooms   map[string]Room // uuid -> Room
}
