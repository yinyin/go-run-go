package main

import (
	"encoding/json"
	"log"

	rungo "github.com/yinyin/go-run-go"
)

func runEnv(cmdGo *rungo.CommandGo) {
	envInfo, err := cmdGo.Env()
	if nil != err {
		log.Printf("failed on invoke env: %v", err)
		return
	}
	buf, err := json.MarshalIndent(envInfo, "", "  ")
	if nil != err {
		log.Printf("failed on marshal (%v): %v", envInfo, err)
		return
	}
	log.Printf("environment information: %s", string(buf))
}
