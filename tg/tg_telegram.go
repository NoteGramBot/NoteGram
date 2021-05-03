/*
 * Notegram / tg
 * Paquete de interfaz con Telegram
 */

package tg

import (
	"fmt"
	"log"
	"strconv"

	tg_botapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBotConfig struct {
	Botconfig BotConfig
	Client  *tg_botapi.BotAPI
	Lastupdate int
}

// tg.Startbot()
// Ejecuta el bot a partir de la configuracion


func NewTelegramBot() (TelegramBotConfig,error) {
	bc,err := NewBot()
	return TelegramBotConfig{Botconfig:bc}, err
}


func (bc *TelegramBotConfig) Connect(apikey string) error {
	bot, err := tg_botapi.NewBotAPI(apikey)
	if err == nil {
		bc.Client = bot
		bc.Botconfig.BotName = bot.Self.FirstName
	}
	return err
}

/*
 * GetMessage()
 * Get the latest message from the bot (BLOCKING)
 */

func (bc *TelegramBotConfig) GetMessage() (BotMessage,error) {

	updates, err := bc.Client.GetUpdatesChan(tg_botapi.NewUpdate(bc.Lastupdate))

	var msg = BotMessage{}

	if err != nil {
		log.Print("Error getting bot update channel ", err)
		return BotMessage{}, err
	}
	
	uu := <- updates // blocks to get message
	bc.Lastupdate = uu.UpdateID + 1

	// TODO: Look at IsMessageToMe(uu)

	// By now, only text

	if (uu.Message.Text != "") {
		msg.ContentType 	= "text/plain"
		msg.Content = []byte(uu.Message.Text)
		msg.From = fmt.Sprintf("%d", uu.Message.Chat.ID) // +- user
		return msg, nil
	}

	log.Print("Warn message without text")
	err = fmt.Errorf("Warning. Message %d without text", uu.UpdateID)
	return msg, err

}

func (bc TelegramBotConfig) SendMessage(msg BotMessage) error {
	chatid, _ := strconv.ParseInt(msg.To, 10,64)
	sendmsg := tg_botapi.NewMessage(chatid, string(msg.Content))
	sendmsg.ParseMode = "markdown"
	_, err := bc.Client.Send(sendmsg)
	return err
}

/* Unimplemented stuff */

func (bc TelegramBotConfig) ListNotes() (msglist []BotMessage, err error) {
	return nil, nil
}

func (bc TelegramBotConfig) Disconnect() error {
    return nil
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