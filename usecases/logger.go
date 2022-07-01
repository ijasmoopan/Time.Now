package usecases

import (
	"log"
	"os"
)

func Logger() *os.File {

	file, err := os.OpenFile("result.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	// defer file.Close()

	return file
}



