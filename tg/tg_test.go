/*
 * Notegram / tg
 * Paquete de interfaz con Telegram
 */

package tg

import (
	"fmt"
	"log"
	"testing"
)


func Test_Example_Newbot(t *testing.T) {
	var the_bot, _ = NewInmemoryBot()	
	err := the_bot.Connect("apikey_123")
	if err != nil {
		log.Fatal("Cannot connect to Messaging service ", err)
	}
	// Get 1 message

	msg, err := the_bot.GetMessage()

	if err != nil {
		log.Fatal("Cannot get messages ", err)
	}

	log.Print("Received Message: ", msg)

	// Send 1 message (same one, as echo)

	msg.To = msg.From
	msg.Content = []byte( fmt.Sprintf(">>> %s", msg.Content) )

	err = the_bot.SendMessage(&msg)

	if err != nil {
		log.Fatal("Cannot send message ", err)
	}
	
}

