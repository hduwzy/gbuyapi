#!/usr/bin/env bash

realpath() {
    cur=$(pwd)
    if [ ! -d $1 ]; then
        return
    fi
    cd $1
    r=$(pwd)
    cd $cur
    echo $r
}

find_all_git() {
    root=$(realpath $1)
    if [ ! -d $root ]; then
        return
    fi

    git_list=$(find $root -type d -name ".git")

    for work_dir in $git_list; do
        echo $(dirname $work_dir)
    done
}

for p in $(find_all_git $1); do
    cur=$(pwd)
    cd $p
    git fetch --all
    cd $cur
    fname=$(basename $p)
    tar -czvf "$fname.tar.gz" "$p/.git"
done