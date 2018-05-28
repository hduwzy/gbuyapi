#!/usr/bin/env bash
export GOPATH="/root/go"
export PATH="$PATH:/root/go/bin"

cd $GOPATH/src/gbuyapi
if [ -f .pid ]; then
    cat .pid | xargs kill
    for id in $(pidof gbuyapi); do
        kill $id
    done
fi

#glide up
nohup bee run -downdoc=true -gendoc=true >/dev/null 2>&1 &
echo $! > .pid
exit