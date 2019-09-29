package main

import (
	"encoding/json"
	"log"

	rungo "github.com/yinyin/go-run-go"
)

func runListPackage(cmdGo *rungo.CommandGo, pkgImportPaths ...string) {
	pkgInfos, err := cmdGo.ListPackage(pkgImportPaths...)
	if nil != err {
		log.Printf("failed on invoke list package: %v", err)
		return
	}
	for resultIndex, pkgInfo := range pkgInfos {
		buf, err := json.MarshalIndent(pkgInfo, "", "  ")
		if nil != err {
			log.Printf("failed on marshal (%v): %v", pkgInfo, err)
			return
		}
		log.Printf("package information (%d): %s", resultIndex, string(buf))
	}
}
