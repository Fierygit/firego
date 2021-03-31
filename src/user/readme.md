




### 用户注册时提供邮箱

    name 
    mail

### 登录



    case1： 从未登录

    - 可以直接使用name向opt_cli 获取动态秘钥 （简单）

    - 写一个服务，浏览器 get firego.cn/user/getopt , 服务器把动态秘钥发送到邮箱， 如果成功认证， 记录 ip， 如果没有认证， 不记录 ip（复杂）




    case2： 已经获取过秘钥

    一定时间内不用验证

    