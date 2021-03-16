/*
 * Notegram / CORE
 * Paquete de logica interna del bot
 */

package core

import (
   "Notegram/data"
   "Notegram/tg"
)

type NotegramError struct {
   msg string
}

// Ejemplo: NotegramError.Error("Invalid Api Key")


func CoreHello() string {
   return "Hello from Core package AND " + data.DataHello() + " AND " + tg.TgHello()
}


