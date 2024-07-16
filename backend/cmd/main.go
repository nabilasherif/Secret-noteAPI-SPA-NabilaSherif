package main

import (
	"flag"
	"fmt"

	"github.com/codescalersinternships/Secret-noteAPI-SPA-NabilaSherif/app"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "port num")
	flag.Parse()

	stringport := ":" + fmt.Sprint(port)
	app.NewApp(stringport)
}
