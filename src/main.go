package main

import (
	"Notegram/core"
	"Notegram/data"
	telegram "Notegram/tg"
	"fmt"
        "flag"
)

func main() {

        // Command line arguments -f config_file_path

        configFilePtr := flag.String("f","config.json", "Path to config file")
        flag.Parse()

	// Bootstrap notegram
	// Obtiene datos de configuracion del fichero.
	// OJO: Algunos datos de la configuracion pueden ser sensibles!

	botconfig, err := core.GetConfig(*configFilePtr)

	if err != nil {
		fmt.Println("Configuration summary:")
		fmt.Print(botconfig)
	}

	fmt.Println("main function")
	fmt.Println(core.CoreHello())
	fmt.Println(data.DataHello())
	fmt.Println(telegram.TgHello())

}
