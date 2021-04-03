/*
 * Notegram / data
 * Paquete de interfaz con la base de datos
 */

package data

type DataError struct {
	msg string
}

func (ee *DataError) Error() string {
	return ee.msg
}

func DataHello() string {
	return "Hello from Data package"
}
