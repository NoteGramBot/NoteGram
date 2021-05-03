/*
 * Notegram / tg
 * Paquete de interfaz con Telegram
 */

package tg

import (
	"testing"
)


func Example_Newbot(t *testing.T) {
	var the_bot = Newbot()
	
}

/*
type fake_telegram_api struct {
	Self: { FirstName: "bot first name"}

} */

/*

func fake_StartBot(k string) (struct, err) {
	fmt.Printf("StartBot called")
}

*/

func test_StartBot(t *testing.T) {
	t.Error("Fail fist")
}
