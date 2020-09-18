

## 创建房间的接口

### post请求 /createroom

### 参数

- isOpen: 0 否， 1 是

- roomName： 这个房间的名字

- roomInfo： 房间的介绍

- roomNotice: 房间的公告

- password： 根据 isOpen 获取

```json
{
    "isOpen" : 0,
    "roomName" : "safd",
    "roomInfo" : "fdf",
    "roomNotice" : "Fdf",
    "passworld" : "df"

}
```

### 返回

- ok
- url（/chatroom?uuid=123456） 

- error
- msg

