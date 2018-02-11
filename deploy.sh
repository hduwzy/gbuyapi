#!/usr/bin/env bash

git add .
git commit -m"auto commit"
git push origin master

ssh root@testonly.fun << eeooff
cd "/root/go/src/gbuyapi" && \
if [ -f .pid ]; then \
    cat .pid | xargs kill -9 \
fi && \
echo 111112222 > .pid
eeooff