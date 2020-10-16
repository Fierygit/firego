###
 # @Author: Firefly
 # @Date: 2020-10-12 16:15:38
 # @Descripttion: 
 # @LastEditTime: 2020-10-16 12:14:00
### 

# 接受所有参数作为 commit 的描述

info=$*
if [ ! $info ]; then 
    echo please input commit info
else 
    echo start to pull from origin main
    git pull
    # how to deal when occur error？
    echo start to push to origin main
    git add .
    git commit -m $info
    git push origin main
fi