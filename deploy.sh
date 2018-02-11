#!/usr/bin/env bash

git add .
git commit -m"auto commit"
git push origin master

ssh root@testonly.fun >/dev/null 2>&1 << eeooff

echo $! > .pid
eeooff