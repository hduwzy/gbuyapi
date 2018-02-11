#!/usr/bin/env bash

cd "/root/go/src/gbuyapi"
if [ -f .pid ]; then
    cat .pid | xargs kill
fi

glide up
nohup bee run -downdoc=true -gendoc=true >/dev/null 2>&1 &
echo $$ > .pid
exit