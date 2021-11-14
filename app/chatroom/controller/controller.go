package chatroom

type ChatRoomController struct {
	rooms *ChatRoom
}

func NewChatRoomController() ChatRoomController {
	return ChatRoomController{
		rooms: InitChatRoom(),
	}
}
