#!/usr/bin/env bash

git add .
git commit -m"auto commit"
git push origin master

ssh root@testonly.fun << eeooff
    cd /root/go/src/gbuyapi
    git pull origin master
eeooff

ssh root@testonly.fun "./deploy_remote.sh"