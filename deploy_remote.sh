#!/usr/bin/env bash
export GOPATH="/root/go"
cd $GOPATH/src/gbuyapi
if [ -f .pid ]; then
    cat .pid | xargs kill
fi

glide up
bee run -downdoc=true -gendoc=true
echo $! > .pid
exit