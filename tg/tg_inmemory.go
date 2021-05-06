/*
 * Notegram / tg
 * Paquete de interfaz con Telegram
 */

package tg

import (
	"fmt"
	"log"
)

type InmemoryBotConfig struct {
	Botconfig BotConfig
	lastmsg  int32
	maxmsg   int32
}

// tg.Startbot()
// Ejecuta el bot a partir de la configuracion


func NewInmemoryBot() (*InmemoryBotConfig,error) {
	bc,err := NewBot()
	return &InmemoryBotConfig{Botconfig:bc, maxmsg: 10}, err
}

func (bc *InmemoryBotConfig) SetMaxMsgs (maxmsg int32) {
	bc.maxmsg = maxmsg
}

func (bc *InmemoryBotConfig) Connect(apikey string) error {
	return nil
}

/*
 * GetMessage()
 * Get the latest message from the bot (BLOCKING)
 */

func (bc *InmemoryBotConfig) GetMessage() (BotMessage,error) {

	var allmsgs []string = 	[]string{
		"😃 Lorem ipsum dolor sit amet,",
		"🧘🏻‍♂️ consectetur adipiscing elit,",
		"🌍 sed eiusmod tempor incidunt",
		"🍞 ut labore et dolore magna aliqua.",
		"🚗 Ut enim ad minim veniam,",
		"📞 quis nostrud exercitation ullamco",
		"🎉 aboris nisi ut aliquid",
		"♥️ ex ea commodi consequat.",
		"🍆 Quis aute iure reprehenderit",
		"📷 in voluptate velit esse cillum dolore",
		"🏁 eu fugiat nulla pariatur." }

	log.Printf("InmemoryBot - maxmsg=%d , lastmsg = %d", bc.maxmsg, bc.lastmsg)

	if (bc.lastmsg < bc.maxmsg) {
		bc.lastmsg++
		var currmsg = allmsgs[bc.lastmsg% int32(len(allmsgs))]

		return BotMessage{
			ContentType: "text/plain",
			Content:     []byte(currmsg),
			From:        "123",
			To:          "234" }, nil

	} else {
		// disconnected
		return BotMessage{}, fmt.Errorf("InmemoryBOT ERROR: Superado limite max mensajes enviados (%d)", bc.maxmsg)
	}

}

func (bc *InmemoryBotConfig) SendMessage(msg *BotMessage) error {
	return nil
}

/* Unimplemented stuff */


func (bc *InmemoryBotConfig) ListNotes() (msglist []BotMessage, err error) {
	return nil, nil
}

func (bc *InmemoryBotConfig) Disconnect() error {
    return nil
}
