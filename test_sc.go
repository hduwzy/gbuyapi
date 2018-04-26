package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.Init()
	l.PushBack(100)
	l.PushBack(200)
	l.PushBack(300)

	head := l.Front()
	for ; head != nil; head = head.Next() {
		fmt.Println(head.Value)
	}
}
