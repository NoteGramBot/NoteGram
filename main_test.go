package main

import (
	"Notegram/core"
	"Notegram/data"
	"Notegram/tg"
	"fmt"
	"log"
	"testing"
)

func Test_BotMain(t *testing.T) {

	var expected_msgs uint32 = 4

	dbclient := data.NotegramStorage{Storage: data.NewBackendInMemory() }
	botclient, _ := tg.NewInmemoryBot()
	botclient.SetMaxMsgs(int32(expected_msgs))

	log.Printf("Testing botmain -> Botclient %+v\n", botclient)

	nc := core.NotegramConfig{
		Secret:       "secret (for testing)",
		Dbhost:       "mongodb host",
		Dbport:       12345,
		Dbuser:       "user",
		Dbpass:       "a password like no other",
		Dbase:        "notegram database",
		Dbcollection: "notegram collection",
		Loglevel:     "muy ruidoso - todavía mas que debug",
	}

	num_msgs := BotMain(nc, botclient, dbclient)

	log.Printf("Botmain: procesados %d mensajes", num_msgs)

	if num_msgs != int32(expected_msgs) {
		errorstr := fmt.Sprintf("No coinciden mensajes recibidos(%d) y esperados (%d)", num_msgs, expected_msgs)
		log.Print(errorstr)
		t.Error(errorstr)
	}


}