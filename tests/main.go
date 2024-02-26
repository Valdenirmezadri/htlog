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

	htl.Log().Info("teste 23")
	htl.Log().Info("teste 24")

}
