/*
 * Notegram / data_memory
 * Interfaz de almacenamiento de BBDD en memoria para testing.
 * Sólo tenemos una nota en memoria (y de un único usuario).
 */

package data

import (
	"Notegram/core"
	"fmt"
)

type DateadorInMemory struct {
	Numnotas int   // numnotas
	Nota     Notes // solo tenemos una nota
	Dateador
}

// Constructor

func NewBackendInMemory() Dateador {
	var dd DateadorInMemory
	dd.Numnotas = 0
	return &dd
}

// Interfaz dateador

func (d *DateadorInMemory) ConnectToDatabase(c core.NotegramConfig) (Dateador, error) {
	d.Numnotas = 0
	return d, nil
}

func (d *DateadorInMemory) Disconnect() {
	d.Numnotas = 0
}

func (d *DateadorInMemory) GetNotas(user string) ([]Notes, error) {
	var retval []Notes

	if d.Numnotas > 0 && user == d.Nota.User {
		retval = append(retval, d.Nota)
		fmt.Print(len(retval))
		fmt.Printf("GetNotas() returns -> %v\n", d.Nota)
		fmt.Printf("Nota (interna) = %v, \n", d.Nota)
	}

	return retval, nil
}

func (d *DateadorInMemory) WriteNota(nota Notes) error {
	d.Numnotas = 1
	d.Nota = nota
	fmt.Printf("Nota (interna) = %v, \n", d.Nota)

	return nil
}

func (dt *DateadorInMemory) DeleteNotaByID(id string) error {
	if dt.Nota.Id == id {
		dt.Numnotas = 0
	}
	return nil
}
