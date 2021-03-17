package main

import (
	"fmt"

	tg_botapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	secret := "*** API KEY QUE NOS HA DADO THE BOTFATHER ***"
	fmt.Println("Using ApiKey: " + secret)

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
