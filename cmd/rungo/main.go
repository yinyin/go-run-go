package main

import (
	"log"
	"os"

	rungo "github.com/yinyin/go-run-go"
)

const rungoCmdUsage = `Argument: [Action] [Options...]

Action:

  * version
  * env
  * list-pkg [PKG_IMPORT_PATH...]
  * list-mod [MOD_IMPORT_PATH...]

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
	case "env":
		runEnv(&cmdGo)
	case "list-pkg":
		runListPackage(&cmdGo, cmdOptions...)
	case "list-mod":
		runListModule(&cmdGo, cmdOptions...)
	case "import-depth":
		runImportDepth(&cmdGo, cmdOptions...)
	}
	return
}
