/*
 * Notegram / data
 * Paquete de interfaz con la base de datos
 */

package data

import (
	"Notegram/core"
	"log"
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

type NotegramStorage struct {
	Storage Dateador
}

// Interfaz dateador para hacer inyección de dependencias
type Dateador interface {
	ConnectToDatabase(core.NotegramConfig) (Dateador, error)
	Disconnect()
	GetNotas(userid string) ([]Notes, error)
	WriteNota(nota Notes) error
	DeleteNotaByID(id string) error
}

func (ee *DataError) Error() string {
	return ee.msg
}

func (ns NotegramStorage) ConnectToDatabase(config core.NotegramConfig) (Dateador, error) {
	return ns.Storage.ConnectToDatabase(config)
}

func (dt NotegramStorage) Disconnect() {
	dt.Storage.Disconnect()
}

func (dt NotegramStorage) GetNotas(user string) ([]Notes, error) {
	result, err := dt.Storage.GetNotas(user)

	if err != nil {
		log.Print("GetNotas: No se pudo realizar la operación ", err)
		return nil, err
	}

	return result, err
}

func (dt NotegramStorage) WriteNota(nota Notes) error {
	err := dt.Storage.WriteNota(nota)
	if err != nil {
		log.Print("WriteNota: no se pudo escribir la nota ", err)
	}
	return err
}

func (dt NotegramStorage) DeleteNotaByID(id string) error {
	err := dt.Storage.DeleteNotaByID(id)
	if err != nil {
		log.Printf("DeleteNota: no se pudo borrar la nota con id=%s %s", id, err)
	}
	return err
}
