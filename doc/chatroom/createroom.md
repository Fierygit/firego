

## 创建房间的接口

### post请求 /createroom

### 参数

- isOpen: 0 否， 1 是

- roomname： 这个房间的名字

- roominfo： 房间的介绍

- roomnotice: 房间的公告

- passworld： 根据 isOpen 获取

### 返回

- ok
- url（/chatroom?uuid=123456） 

- error
- msg
- code

1、 
2、