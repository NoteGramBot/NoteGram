package data

import (
	"Notegram/core"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

var cfg = core.NotegramConfig{
	Secret:       "somesecret",
	Dbhost:       "127.0.0.1",
	Dbport:       27017,
	Dbuser:       "scott",
	Dbpass:       "tiger",
	Dbase:        "notegram",
	Dbcollection: "notas",
	Loglevel:     "Debug",
}

func TestConecta(t *testing.T) {

	var dat = NotegramStorage{Storage: NewBackendInMemory()}
	conn, err := dat.ConnectToDatabase(cfg)
	fmt.Println("conn: ", conn, " err: ", err)

	// Para este test tenemos una bbdd en local con esas credenciales
	// o bien un dateador que nos "mockea" la conexion a mongo etc..

	if err != nil {
		// esto suena a integration test fallido
		t.Error("ConnectToDatabase(", cfg, ") FAILED wirh error ", err)
		t.FailNow()
	}

	conn.Disconnect()

}

/*
 * Comprobamos que funcionan las operaciones de la BBDD_
 * almacenando y recuperando el mismo registro
 */

func Test_WriteReadDelete(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	var testUsername = fmt.Sprintf("testuser_%020d", rand.Int63n(1000000000000000000))

	var newNote = Notes{
		Id:              "0x00000000",
		User:            testUsername,
		Content:         "👏👏👏  HELLO",
		ContentType:     "text/plain",
		ContentEncoding: "utf8",
	}

	dat := NotegramStorage{Storage: NewBackendInMemory()}

	_, err := dat.ConnectToDatabase(cfg)

	defer dat.Disconnect()

	if err != nil {
		// esto suena a integration test fallido
		t.Fatal("TestWrite(", cfg, ") FAILED trying to connect to database", err)
	}

	err = dat.WriteNota(newNote)

	fmt.Println("Writenota() -> err = ", err)
	fmt.Println("Writenota backend ", dat)

	regs, err := dat.GetNotas(testUsername)
	if err != nil {
		t.Error("TestWrite: Fallo al escribir en la BBDD")
	}

	fmt.Printf("TESTNOTAS: GetNotas(%s) = %+v\n", testUsername, regs)

	fmt.Println("TESTNOTAS:\n\terr = ", err, "\n\t regs = ", regs, "len(regs)=", len(regs))

	readNote := regs[len(regs)-1]
	readDocid := readNote.Id

	log.Println("Registro escrito:\n\t", newNote)
	log.Println("Registro leido:\n\t", readNote)

	readNote.Id = "0x00000000"

	if readNote != newNote {
		t.Error("Las notas son diferentes -  escrito:", newNote, " leido: ", regs)
	}

	// Borramos el registro y comprobamos que no existe al leerlo

	err = dat.DeleteNotaByID(readDocid)

	if err != nil {
		t.Errorf("Error al borrar la nota con id: %s", readDocid)
	}

	// Volvemos a buscar el registro

	regs, err = dat.GetNotas(testUsername)

	fmt.Println("TESTNOTAS:\n\terr = ", err, "\n\t regs = ")

	if len(regs) != 0 {
		log.Println("Se ha encontrado un registro que debería estar borrado")
	}

	log.Println("Buscamos despues de borrar. Resultados: err=", err, "regs=", regs)

}
