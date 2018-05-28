package main

import (
	"gbuyapi/services"
	"os"
)

func main() {
	s := services.NewWechatService()
	s.Login(os.Args[1])
}