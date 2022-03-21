package main

import "C"
import (
	shell "goterm/src"
	"log"
)

func main() {
	comm, err := shell.Parses("ls -lha && cat main.go")
	if err != nil {
		log.Fatalln(err)
	}

	shell.Exec_shell(comm)
}
