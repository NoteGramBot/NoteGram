package main

import (
   "fmt"
   "Notegram/core"
   "Notegram/data"
   telegram "Notegram/tg"
)


func main() {

   fmt.Println("main function")
   fmt.Println(core.CoreHello())
   fmt.Println(data.DataHello())
   fmt.Println(telegram.TgHello())

}


