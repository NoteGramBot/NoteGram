package core

import (
	"testing"
)

// Core function testing

func Test_ConfigFileNotFound(t *testing.T) {

	_, err := GetConfig("/this/filename/does/not/exist/even/if/you/are/root/.yeison")

	if err == nil {
		t.Error("No se deberia encontrar un fichero con un nombre tan raro")
	}

}

func Test_InvalidJsonFile(t *testing.T) {

	_, err1 := GetConfig("testdata/malformed_config.json")
	_, err2 := GetConfig("testdata/malformed_config.txt")
	_, err3 := GetConfig("testdata/malformed_config2.json")
	// _, err4 := GetConfig("testdata/valid_config.json")

	errlist := []error{err1, err2, err3}
	for _, ee := range errlist {
		if ee == nil {
			t.Error("Check Kaputt. Deberias tener un error al leer ficheros JSON de configuracion txungos")
		}
	}

}
