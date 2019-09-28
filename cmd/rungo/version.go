package main

import (
	"log"

	rungo "github.com/yinyin/go-run-go"
)

func runVersion(cmdGo *rungo.CommandGo) {
	majorVer, minorVer, outputText, err := cmdGo.Version()
	log.Printf("version: %d, %d; err: %v; output-text: %v", majorVer, minorVer, err, outputText)
}
