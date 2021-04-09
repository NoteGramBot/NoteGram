/*
 * Notegram / data
 * Paquete de interfaz con la base de datos
 */

package data

import (
	"Notegram/core"
)

type DataError struct {
	msg string
}

type Notes struct {
	Id              string `bson:"-"`
	User            string `bson:"user"`
	Content         string `bson:"content"`
	ContentType     string `bson:"content_type"`
	ContentEncoding string `bson:"content_encoding.ifpresent"`
}

// Interfaz dateador para hacer inyección de dependencias
type Dateador interface {
	ConnectToDatabase(core.NotegramConfig) (Dateador, error)
	Disconnect()
	GetNotas(userid string) ([]Notes, error)
	WriteNota(nota Notes) error
	DeleteNotaByID(id string)
}

func (ee *DataError) Error() string {
	return ee.msg
}
