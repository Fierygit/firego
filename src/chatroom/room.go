package chatroom

import "time"

type Client struct {
	IP            string // 用户的ip
	Name          string // 用户的昵称， 一个要一个， 前段可以随机生成
	Index         int    // 已经拉去消息的索引位置
	IsAlive       bool   // 这个用户时候还在
	LastAliveTime int64  // 上一次在线的时间
}

type Message struct {
	Type     int8        // 消息类型， 0 为普通消息， 1 为提示消息等等
	Info     interface{} // 存储消息
	SendTime time.Time   // 存储这条消息的时间
}

type RoomInfo struct {
	CreateInfo *CreateRoomReq // 创建房间必要的参数
	CreateTime time.Time      // 创建的时间
}

type Room struct {
	RoomInfo  RoomInfo
	ClientNum int                //	在线的人数
	Messages  []*Message         // 消息列表， 索引号按时间排序
	Clients   map[string]*Client //	SessionId -> Client
}

// chatroom 的结构体
// 暂时不支持通过房间名字查询
type ChatRoom struct {
	Rooms map[string]*Room // uuid -> Room
}

// -------------------------- operation -------------------------------------
//--------------------------------Message--------------------------------------
func MakeMessage(mType int8, info interface{}) *Message {
	return &Message{mType, info, time.Now()}
}

//--------------------------------RoomInfo--------------------------------------
func MakeRoomInfo(CreateInfo *CreateRoomReq) *RoomInfo {
	return &RoomInfo{CreateInfo, time.Now()}
}

// --------------------------------Room--------------------------------------
func (room *Room) AddMessage(message *Message) {
	messages := room.Messages
	room.Messages = append(messages, message)
}

// 从 0 开始
func (room *Room) GetMessage(start int, end int) ([]*Message, string) {
	if start > end {
		return nil, ""
	}
	maxLen := len(room.Messages) - 1
	if end > maxLen {
		end = maxLen
	}
	return room.Messages[start:end], ""
}

func (room *Room) AddClient(sessionId string, client *Client) {
	room.Clients[sessionId] = client
}

func (room *Room) RemoveClient(sessionId string) {
	room.Clients[sessionId] = nil
}

func (room *Room) GetClient(sessionId string) *Client {
	return room.Clients[sessionId]
}

func (room *Room) findClientBySessionId(id string) bool {
	return room.Clients[id] != nil
}

//---------------------------chatRoom-------------------------------------
func InitChatRoom() *ChatRoom {
	return &ChatRoom{map[string]*Room{}}
}

// 当已经存在房间了返回false
// todo 冲突，还是重复了？
func (chatRoom *ChatRoom) AddRoom(uuid string, createRoomArg *CreateRoomReq) bool {
	if chatRoom.Rooms[uuid] != nil {
		return false
	}
	room := Room{
		RoomInfo: RoomInfo{
			createRoomArg,
			time.Now()},
		ClientNum: 0,
		Messages:  []*Message{},
		Clients:   map[string]*Client{}}
	chatRoom.Rooms[uuid] = &room
	return true
}

func (chatRoom *ChatRoom) GetRoom(uuid string) *Room {
	return chatRoom.Rooms[uuid]
}

func (chatRoom *ChatRoom) RemoveRoom(uuid string) {
	chatRoom.Rooms[uuid] = nil
}

func (chatRoom *ChatRoom) RoomNum() int {
	return len(chatRoom.Rooms)
}
