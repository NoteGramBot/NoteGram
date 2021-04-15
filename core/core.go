/*
 * Notegram / CORE
 * Paquete de logica interna del bot
 */

package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type NotegramError struct {
	msg string
}

func (ee *NotegramError) Error() string {
	return ee.msg
}

// notegramConfig - los campos tienen que emprezar en mayuscula para exportarse
// si no, el paquete encoding/json no puede verlos, y no emite ningún error.

type NotegramConfig struct {
	Secret       string // Telegram secret
	Dbhost       string // mongodb hostname/ip
	Dbport       uint   // mongodb port
	Dbuser       string // mongodb username
	Dbpass       string // mongodb passwd
	Dbase        string // mongodb database
	Dbcollection string // mongodb collection
	Loglevel     string // loglevel
}

/*
 * getConfig()
 * Obtiene la información de configuracion del bot en forma de map (como el dict de python, o hash de Perl)
 * - Telegram Secret
 * - Host/port/pass/index al que escribir en mongodb
 * - Flag debug
 */

func GetConfig(jsonfilename string) (NotegramConfig, error) {

	var conf NotegramConfig
	err := error(nil)

	jsconfig, ee := ioutil.ReadFile(jsonfilename)

	if ee != nil {
		log.Printf("No se ha podido leer el fichero %s", jsonfilename)
		return conf, ee
	}

	err = json.Unmarshal(jsconfig, &conf)

	if err != nil {
		fmt.Println("error reading json file")
	}

	fmt.Println(conf)

	return conf, err // the rust borrow checker would love this.

}
