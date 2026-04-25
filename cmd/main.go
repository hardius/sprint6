package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	myLogger := log.New(os.Stdout, "Logger: ", log.LstdFlags)

	myServer := server.NewRouter(myLogger)
	if err := myServer.Server.ListenAndServe(); err != nil {
		fmt.Println("ой")
		myLogger.Fatal(err)
	}
}
