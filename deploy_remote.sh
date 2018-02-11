#!/usr/bin/env bash
export GOPATH="/root/go/src/gbuyapi"
cd $GOPATH
if [ -f .pid ]; then
    cat .pid | xargs kill
fi

glide up
nohup bee run -downdoc=true -gendoc=true >/dev/null 2>&1 &
echo $$ > .pid
exit