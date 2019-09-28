package main

import (
	"./src/bindShell"
	//"./src/fileServer"
)

func main() {
	//port := flag.String("p", "8100", "port to serve on")
	//directory := flag.String("d", ".", "the directory of static file to host")
	//flag.Parse()
	bindShell.Run()
	//fileServer.Run("443", ".")

}
