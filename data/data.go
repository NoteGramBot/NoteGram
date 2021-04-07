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
	// unused // ctx          context.Context
}

type Notes struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	User            string             `bson:"user"`
	Content         string             `bson:"content"`
	ContentType     string             `bson:"content_type"`
	ContentEncoding string             `bson:"content_encoding.ifpresent"`
}

// ConnectToDatabase: Conecta y hace ping al servidor
// config.Connect() <-- Esta pidiendo una interfaz como esta ;-)
func ConnectToDatabase(config core.NotegramConfig) (NotegramConnection, error) {

	var dc NotegramConnection

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() 

	var connURI string = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		config.Dbuser,
		config.Dbpass,
		config.Dbhost,
		config.Dbport,
		config.Dbase)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connURI))

	// Completamos la estructura para llamadas siguientes
	dc.mongocli = client
	dc.dbcollection = config.Dbcollection
	dc.database = config.Dbase

	err = client.Ping(ctx, nil)

	return dc, err
}

func (conn NotegramConnection) Disconnect() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() 

	conn.mongocli.Disconnect(ctx)
	conn.mongocli = nil
}

// Interfaz CRUD

func (conn NotegramConnection) GetNotas(userid string) ([]Notes, error) {
	var notas []Notes
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() 

	db := conn.mongocli.Database(conn.database)
	coll := db.Collection(conn.dbcollection)
	cursor, err := coll.Find(ctx, bson.M{"user": userid})
	err = cursor.All(ctx, &notas)

	return notas, err
}

// WriteNota(nota)
// Escribe una nota en la BBDD mongodb.
func (conn NotegramConnection) WriteNota(nota Notes) error {
	db := conn.mongocli.Database(conn.database)
	coll := db.Collection(conn.dbcollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() 
	_, err := coll.InsertOne(ctx, &nota)
	return err
}

// DeleteNotaById(id)
// Borra la nota con el ID de mensaje pasada como argumento
func (conn NotegramConnection) DeleteNotaByID(id primitive.ObjectID) error {
	db := conn.mongocli.Database(conn.database)
	coll := db.Collection(conn.dbcollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() 
	_, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (ee *DataError) Error() string {
	return ee.msg
}
