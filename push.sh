
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