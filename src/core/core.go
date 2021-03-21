/*
 * Notegram / CORE
 * Paquete de logica interna del bot
 */

package core

import (
	"Notegram/data"
	"Notegram/tg"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Secret       string
	Dbhost       string
	Dbport       uint
	Dbuser       string
	Dbpass       string
	Dbcollection string
	Loglevel     string
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
		fmt.Println("No se ha podido leer el fichero config.json\n")
		return conf, ee
	}

	err = json.Unmarshal(jsconfig, &conf)

	if err != nil {
		fmt.Println("error reading json file")
	}

	fmt.Println(conf)

	return conf, err // the rust borrow checker would love this.

}

func CoreHello() string {
	return "Hello from Core package AND " + data.DataHello() + " AND " + tg.TgHello()
}
