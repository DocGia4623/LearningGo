package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

func main() {

	//add log
	log.SetFormatter(&ecslogrus.Formatter{})
	log.SetLevel(log.TraceLevel)

	logfilePath := "/logs/out.log"
	file, err := os.OpenFile(logfilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	}
	defer file.Close()
	fmt.Print("Start device")
	log.Info("Start Service")

	server := &http.Server{
		Addr: ":8080",
		//handle routes
		ReadTimeout:   10 * time.Second,
		WriteTimeout:  10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	if server_err != nil {
		panic(server_err)
	}
	
}


