/*
 * Notegram / data
 * Paquete de interfaz con la base de datos
 */

package data

import (
	"Notegram/core"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataError struct {
	msg string
}

type NotegramConnection struct {
	mongocli     *mongo.Client // Cliente de la base de datos
	database     string        // Where data is stored
	dbcollection string        // Collection en la que buscamos
	tguser       string        // Usuario de telegram
	lastnote     string        // cache
	ctx          context.Context
}

type Notes struct {
	Id              primitive.ObjectID `bson:"_id"`
	User            string             `bson:user`
	Content         string             `bson:"content"`
	ContentType     string             `bson:"content_type"`
	ContentEncoding string             `bson:"content_encoding.ifpresent`
}

// ConnectToDatabase: Conecta y hace ping al servidor
// config.Connect() <-- Esta pidiendo esto ;-)
func ConnectToDatabase(config core.NotegramConfig) (NotegramConnection, error) {

	var dc NotegramConnection

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// was: ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	var connURI string = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		config.Dbuser,
		config.Dbpass,
		config.Dbhost,
		config.Dbport,
		config.Dbase)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connURI))

	// Necesitamos esto para hacer un find
	dc.mongocli = client
	dc.dbcollection = config.Dbcollection
	dc.database = config.Dbase
	dc.ctx = ctx // Necesitamos el context para llamar a mongodb

	err = client.Ping(ctx, nil)

	return dc, err
}

func (conn NotegramConnection) Disconnect() {
	conn.mongocli.Disconnect(conn.ctx)
	conn.mongocli = nil
	conn.ctx = nil
}

// Interfaz CRUD

func (conn NotegramConnection) GetNotas(userid string) []Notes {

	// Hay notas ???

	var notas []Notes

	db := conn.mongocli.Database(conn.database)
	coll := db.Collection(conn.dbcollection)

	err := coll.FindOne(conn.ctx, bson.M{"_id": userid}).Decode(&notas)

	fmt.Println(err)

	return notas

}

func WriteNotaUser(config NotegramConnection, username string, nota string) error {

	// Primero obtenemos todas las notas
	// Si no hay notas, escribimos una entrada en la collection para ese usuario con una nota
	// Si hay notas escribimos una más.

	panic("Not implemented")

}

func (ee *DataError) Error() string {
	return ee.msg
}

// deprecated
func DataHello() string {
	return "Hello from Data package"
}
