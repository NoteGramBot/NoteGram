/*
 * Notegram / tg
 * Paquete de interfaz con Telegram
 */

package tg

import (
	"fmt"
	"log"

	tg_botapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBotConfig struct {
	Botconfig BotConfig
	Client  *tg_botapi.BotAPI
	lastupdate int
}

// tg.Startbot()
// Ejecuta el bot a partir de la configuracion


func NewTelegramBot() (TelegramBotConfig,error) {
	bc,err := NewBot()
	return TelegramBotConfig{Botconfig:bc}, err
}


func (bc *TelegramBotConfig) Connect(apikey string) {

	bot, err := tg_botapi.NewBotAPI(apikey)

	if err == nil {
		bc.Client = bot
		bc.Botconfig.BotName = bot.Self.FirstName
	}
}

/*
 * GetMessage()
 * Get the latest message from the bot (BLOCKING)
 */

func (bc *TelegramBotConfig) GetMessage() (BotMessage,error) {

	updates, err := bc.Client.GetUpdatesChan(tg_botapi.NewUpdate(bc.lastupdate))

	var msg = BotMessage{}

	if err != nil {
		log.Print("Error getting bot update channel ", err)
		return BotMessage{}, err
	}
	
	uu := <- updates // blocks to get message
	bc.lastupdate = uu.UpdateID + 1

	// By now, only text

	if (uu.Message.Text != "") {
		msg.ContentType 	= "text/plain"
		msg.Content = []byte(uu.Message.Text)
		return msg, nil
	}

	log.Print("Warn message without text")
	err = fmt.Errorf("Warning. Message %d without text", uu.UpdateID)
	return msg, err

}

/*
func Do_Bot(secret string) {

	bot, err := tg_botapi.NewBotAPI(secret)

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

*/