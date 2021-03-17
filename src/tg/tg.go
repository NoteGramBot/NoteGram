/*
 * Notegram / tg
 * Paquete de interfaz con Telegram
 */

package tg

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


