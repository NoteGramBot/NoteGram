/*
 * Notegram / tg
 * Paquete de interfaz con Telegram
 */

package tg

import (
    "fmt"
    tg_botapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Nota: Habría que llevar un conteo de tipos de error etc...
type TelegramError struct {
   msg string
}

func (ee *TelegramError) Error() string {
    return ee.msg
}

func TgHello() string {
   return "Hello from TG Package"
}

// tg.Startbot()
// Ejecuta el bot a partir de la configuracion

func StartBot(secret string) {

   bot, err := tg_botapi.NewBotAPI(secret)

   if err != nil {
	fmt.Print("Err: " + err.Error())
   }

   fmt.Println("Bot Name: " + bot.Self.FirstName)

   updates, err := bot.GetUpdatesChan(tg_botapi.NewUpdate(0))

   for uu := range updates {

      if uu.Message != nil {
		fmt.Println("Update: " + uu.Message.Text)
		// Devuelve el mensaje (echo server)
		sendmsg := tg_botapi.NewMessage(uu.Message.Chat.ID, "> "+uu.Message.Text)
		sendmsg.ParseMode = "markdown"
		bot.Send(sendmsg)
	}

   }

}
