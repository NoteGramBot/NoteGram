package main

import (
	"Notegram/core"
	telegram "Notegram/tg"
	"flag"
	"fmt"
)

func main() {

	// Command line arguments -f config_file_path

	configFilePtr := flag.String("f", "config.json", "Path to config file")
	flag.Parse()

	// Bootstrap notegram
	// Obtiene datos de configuracion del fichero.
	// OJO: Algunos datos de la configuracion pueden ser sensibles!

	botconfig, err := core.GetConfig(*configFilePtr)

	if err != nil {
		fmt.Println("Configuration summary:")
		fmt.Print(botconfig)
	}

	telegram.StartBot(botconfig.Secret)

	fmt.Println("main function")
	fmt.Println(core.CoreHello())
	fmt.Println(telegram.TgHello())

}
