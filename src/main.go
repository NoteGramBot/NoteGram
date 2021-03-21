package main

import (
	"Notegram/core"
	"Notegram/data"
	telegram "Notegram/tg"
	"fmt"
)

func main() {

	// Bootstrap notegram
	// Obtiene datos de configuracion del fichero.
	// OJO: Algunos datos de la configuracion pueden ser sensibles!

	botconfig, err := core.GetConfig("config.json")

	if err != nil {
		fmt.Println("Configuration summary:\n")
		fmt.Print(botconfig)
	}

	fmt.Println("main function")
	fmt.Println(core.CoreHello())
	fmt.Println(data.DataHello())
	fmt.Println(telegram.TgHello())

}
