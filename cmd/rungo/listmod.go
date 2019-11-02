package main

import (
	"encoding/json"
	"log"

	rungo "github.com/yinyin/go-run-go"
)

func runListModule(cmdGo *rungo.CommandGo, modImportPaths ...string) {
	modInfos, err := cmdGo.ListModule(modImportPaths...)
	if nil != err {
		log.Printf("failed on invoke list module: %v", err)
		return
	}
	for resultIndex, modInfo := range modInfos {
		buf, err := json.MarshalIndent(modInfo, "", "  ")
		if nil != err {
			log.Printf("failed on marshal (%v): %v", modInfo, err)
			return
		}
		log.Printf("module information (%d): %s", resultIndex, string(buf))
	}
}
