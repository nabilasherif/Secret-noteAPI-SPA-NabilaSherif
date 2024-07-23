package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/codescalersinternships/Secret-noteAPI-SPA-NabilaSherif/app"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "port num")
	var dbPath string
	flag.StringVar(&dbPath, "d", "", "database filepath")
	var secretKey string
	flag.StringVar(&secretKey, "k", "", "secret key for the tokens")
	flag.Parse()

	stringport := ":" + fmt.Sprint(port)
	appInstance, err := app.NewApp(stringport, dbPath, secretKey)
	if err != nil {
		log.Fatal(err)
	}

	if err := appInstance.Run(stringport); err != nil {
		log.Fatal(err)
	}
}
