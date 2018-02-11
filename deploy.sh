#!/usr/bin/env bash

git add .
git commit -m"auto commit"
git push origin master

ssh root@testonly.fun << eeooff
cd /root/go/src/gbuyapi
if [ -f ".pid" ]; then
    cat .pid | xargs kill -9
fi
git pull origin master
glide up
nohup bee run -downdoc=true -gendoc=true >/dev/null 2>&1
echo $! > .pid
eeooff