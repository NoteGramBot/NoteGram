/*
 * Notegram / tg
 * Paquete de interfaz con Telegram
 */

package tg

type TelegramError struct {
   msg string
}

// Ejemplo: TelegramError.Error("Invalid Api Key")


func TgHello() string {
   return "Hello from TG Package"
}


