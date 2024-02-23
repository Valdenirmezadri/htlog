package main

import (
	"log"

	htl "github.com/Valdenirmezadri/htlog"
)

func main() {
	err := htl.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer htl.Stop()

	htl.Warning("teste 2")
	htl.Warning("teste 3")

}
