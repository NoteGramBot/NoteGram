package main

import (
	"Notegram/core"
	"Notegram/data"
	"Notegram/tg"
	"flag"
	"log"
)

func main() {

	// Command line arguments -f config_file_path

	configfile := flag.String("f", "config.json", "Path to config file")
	flag.Parse()

	// Bootstrap notegram
	// Obtiene datos de configuracion del fichero.
	// OJO: Algunos datos de la configuracion pueden ser sensibles!

	botconfig, err := core.GetConfig(*configfile)

	if err != nil {
		log.Printf("Error reading config file %s", *configfile)
		log.Fatalf("Configuration read: %+v", botconfig )
	}

	// BotMain is a separate function to be able to test it ;-)

	botclient, err := tg.NewTelegramBot() // new _telegram_ bot, not in-memory bot

	if err != nil {
		log.Fatal("Cannot create bot client")
	}

	dbclient := data.NotegramStorage{Storage: data.NewBackendMongodb()}


	BotMain(botconfig, botclient, dbclient)

}

func BotMain (conf core.NotegramConfig, botclient tg.BotInterface, db data.NotegramStorage) {

	// Setup the bot

	err := botclient.Connect(conf.Secret)
	if err != nil {
		log.Fatal("Cannot connect to Messaging service ", err)
	}

	_, err = db.ConnectToDatabase(conf)
	defer db.Disconnect()

	if err != nil {
		log.Fatal("Cannot connect to Database ", err)
	}

	for 1 > 0 {
		// Blocks until we get a message!
		recvmsg, err := botclient.GetMessage()
		if err != nil {
			// maybe this timed out
			log.Fatal("Cannot get messages ", err)
		}

		var usernote = data.Notes {
			Id:              "0x00000000", // ya asignará una el backend
			User:            recvmsg.From,
			Content:         string(recvmsg.Content),
			ContentType:     recvmsg.ContentType,
			ContentEncoding: "utf8",	
		}

		// Store the message in the Database
		db.WriteNota(usernote)
	}

}
