/*
 * Notegram / tg
 * Paquete de interfaz con Telegram
 */

package tg

// Nota: Habría que llevar un conteo de tipos de error etc...
type TelegramError struct {
	msg string
}

type BotConfig struct {
	Apikey	string	// Api key encoded as a string
	BotName	string  // Name to show
}

type BotMessage struct {
	ContentType string
	Content []byte
	From string // ChatID as string
	To string // ChatID should go here
}

type BotInterface interface {
	Connect(string) error
	GetMessage() (*BotMessage, error) 
	SendMessage(*BotMessage) error 
	ListNotes() ([]BotMessage, error)
	Disconnect(*BotConfig) error
	// GetNote(uint64) (*BotMessage, error)
}

func (ee *TelegramError) Error() string {
	return ee.msg
}

/*
 * NetBot(apikey string)
 * Creates new bot structure
 * Returns valid *Botconfig on success, error
 */
func NewBot() (BotConfig, error) {	
	return BotConfig{}, nil
}
