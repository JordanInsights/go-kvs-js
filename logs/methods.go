package logs

import (
	"fmt"
	"log"
	"os"
)

var (
	ErrorLogger   *log.Logger
	RequestLogger *log.Logger
)

func Init() {
	errorLog, err := os.OpenFile("errors.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	requestLog, err := os.OpenFile("htaccess.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	ErrorLogger = log.New(errorLog, "", log.Ldate|log.Ltime|log.Lshortfile)
	RequestLogger = log.New(requestLog, "", log.Ldate|log.Ltime)
}

func LogRequest(remoteAddr string, method string, path string) {
	f, err := os.OpenFile("htaccess.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	fmt.Println(path)

	log.SetOutput(f)
	log.Printf("%s %s %s", remoteAddr, method, path)
}
