/*
 * Notegram / data
 * Paquete de interfaz con la base de datos
 */

package data

type DataError struct {
   msg string
}

// Ejemplo: DataError.Error("No se puede conectar a la BBDD")


func DataHello() string {
   return "Hello from Data package"
}


