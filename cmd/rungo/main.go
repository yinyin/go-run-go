package main

import (
	"log"
	"os"

	rungo "github.com/yinyin/go-run-go"
)

const rungoCmdUsage = `Argument: [Action] [Options...]

Action: version

`

func main() {
	if len(os.Args) < 2 {
		log.Print(rungoCmdUsage)
		return
	}
	cmdAction := os.Args[1]
	cmdOptions := os.Args[2:]
	log.Printf("Command: action=%v, options=%v", cmdAction, cmdOptions)
	cmdGo := rungo.CommandGo{}
	switch cmdAction {
	case "version":
		runVersion(&cmdGo)
	}
	return
}
