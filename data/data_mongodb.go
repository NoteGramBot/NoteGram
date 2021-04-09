/*
 * Notegram / data
 * Implementa la interfaz con MongoDB
 */

package data

import (
	"Notegram/core"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Implementamos interfaz Dateador para almacenamiento de notas
//
// type Dateador interface {
//	ConnectToDatabase(core.NotegramConfig) (Dateador, error)
//	Disconnect()
//	GetNotas(userid string) ([]Notes, error)
//	WriteNota(nota Notes) error
//	DeleteNotaByID(id string)
// }

type NotegramConnection struct {
	mongocli     *mongo.Client // Cliente de la base de datos
	database     string        // Where data is stored
	dbcollection string        // Collection en la que buscamos
	tguser       string        // Usuario de telegram
	lastnote     string        // cache
	// unused // ctx          context.Context
}

type NotesMongo struct {
	mongoId primitive.ObjectID `bson:"_id,omitempty"`
	Notes
}

func (n Notes) to_mongo() NotesMongo {
	var nm NotesMongo
	nm.Notes = n
	fmt.Print("Calling: ", n)

	oid, _ := primitive.ObjectIDFromHex(n.Id)
	nm.mongoId = oid
	fmt.Print(nm)

	return nm // el compilador asigna nm en el heap, no en la pila???
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
	var retval []Notes

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("DEBUG GetNotas(%s)\n", userid)

	db := conn.mongocli.Database(conn.database)
	coll := db.Collection(conn.dbcollection)

	filter := bson.M{"user": userid}

	log.Printf("GetNotas - filter= %v\n", filter)
	cursor, err := coll.Find(ctx, filter)

	log.Printf("*****\ncoll.Find() ->\n\t cursor = r%+v, err %+v \n", cursor, err)

	if err != nil {
		log.Println("**************\n**********\n coll.Find() returned error -> ", err)
	}
	err = cursor.All(ctx, &retval)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Getnotas(Userid = %s):\n\t Notasmongo = %+v", userid, retval)

	return retval, err
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
func (conn NotegramConnection) DeleteNotaByID(id string) error {

	oid, err := primitive.ObjectIDFromHex(id)
	if err == nil {
		// Este error debería estar en el inventario
		log.Print("Invalid id (must be 12byte hex string)")
		return err
	}

	db := conn.mongocli.Database(conn.database)
	coll := db.Collection(conn.dbcollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := coll.DeleteOne(ctx, bson.M{"_id": oid})

	fmt.Print("Resultado borrado:", res)

	return err
}
