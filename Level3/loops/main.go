package main

import (
  "fmt"
)

func main() {
  for  i:=5;i<10;i++{
    fmt.Println(i)
  }
  isAlive := false
  for !isAlive {
    isAlive = util.PingChatServer()
  }
  for {
    if util.PingChatServer() {
      break
    }
  }
}
