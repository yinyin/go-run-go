package main

import (
	"errors"
	"log"
	"strings"

	rungo "github.com/yinyin/go-run-go"
)

type packageImportReference struct {
	ImportPath   string
	PackageInfo  *rungo.Package
	ImportsOther []*packageImportReference // import other package
	BeImports    []*packageImportReference // import by other package
	DepthValue   int
}

func (r *packageImportReference) setDepthValue(v int) {
	if v <= r.DepthValue {
		return
	}
	r.DepthValue = v
	for _, impOtherRef := range r.ImportsOther {
		impOtherRef.setDepthValue(v + 1)
	}
}

func (r *packageImportReference) addImportsOther(ref *packageImportReference) {
	for _, impOtherRef := range r.ImportsOther {
		if (impOtherRef == ref) || (impOtherRef.ImportPath == ref.ImportPath) {
			return
		}
	}
	r.ImportsOther = append(r.ImportsOther, ref)
	ref.addBeImports(r)
	ref.setDepthValue(r.DepthValue + 1)
}

func (r *packageImportReference) addBeImports(ref *packageImportReference) {
	for _, beImpRef := range r.BeImports {
		if (beImpRef == ref) || (beImpRef.ImportPath == ref.ImportPath) {
			return
		}
	}
	r.BeImports = append(r.BeImports, ref)
}

type packageImportCollector struct {
	ImportRefs     map[string]*packageImportReference
	RootImportPath string
	CmdGo          *rungo.CommandGo
}

func newPackageImportCollector(cmdGo *rungo.CommandGo, rootImportPath string) (result *packageImportCollector) {
	return &packageImportCollector{
		ImportRefs:     make(map[string]*packageImportReference),
		RootImportPath: rootImportPath,
		CmdGo:          cmdGo,
	}
}

func (c *packageImportCollector) fetchPackageInfo(importPath string) (pkgRef *packageImportReference, err error) {
	pkgInfos, err := c.CmdGo.ListPackage(importPath)
	if nil != err {
		log.Printf("failed on fetch package information for %v: %v", importPath, err)
		return
	}
	if len(pkgInfos) != 1 {
		log.Print("unexpect result package information count for %v: %d", importPath, len(pkgInfos))
		err = errors.New("unexpect result count")
		return
	}
	pkgInfo := pkgInfos[0]
	pkgRef = c.addReference(importPath, pkgInfo)
	return
}

func (c *packageImportCollector) addReference(importPath string, pkgInfo *rungo.Package) (pkgRef *packageImportReference) {
	pkgRef, ok := c.ImportRefs[importPath]
	if ok {
		return
	}
	c.ImportRefs[importPath] = nil
	pkgRef = &packageImportReference{
		ImportPath:  importPath,
		PackageInfo: pkgInfo,
		DepthValue:  0,
	}
	for _, impIntoPath := range pkgInfo.Imports {
		if !strings.HasPrefix(impIntoPath, c.RootImportPath) {
			log.Printf("skip processing import: %s", impIntoPath)
			continue
		}
		var err error
		impRef, ok := c.ImportRefs[impIntoPath]
		if !ok {
			if impRef, err = c.fetchPackageInfo(impIntoPath); nil != err {
				log.Printf("failed to fetch imported package information %v: %v", impIntoPath, err)
				return nil
			}
		}
		if nil == impRef {
			log.Printf("failed to get imported package reference: %v", impIntoPath)
			return nil
		}
		pkgRef.addImportsOther(impRef)
	}
	c.ImportRefs[importPath] = pkgRef
	return
}

func (c *packageImportCollector) withImportOrder() (result []*packageImportReference, maxDepth int) {
	maxDepth = -1
	for _, pkgRef := range c.ImportRefs {
		if pkgRef.DepthValue > maxDepth {
			maxDepth = pkgRef.DepthValue
		}
	}
	for depth := maxDepth; depth >= 0; depth-- {
		for _, pkgRef := range c.ImportRefs {
			if pkgRef.DepthValue == depth {
				result = append(result, pkgRef)
			}
		}
	}
	return
}

func fetchRootImportPath(cmdGo *rungo.CommandGo) (rootImportPath string) {
	rootPkgInfos, err := cmdGo.ListPackage()
	if nil != err {
		log.Printf("failed on invoke list package for root path: %v", err)
		return
	}
	if len(rootPkgInfos) != 1 {
		log.Printf("unexpect package information: %v", rootPkgInfos)
		return
	}
	rootImportPath = rootPkgInfos[0].ImportPath
	return
}

func runImportDepth(cmdGo *rungo.CommandGo, mainImportPaths ...string) {
	rootImportPath := fetchRootImportPath(cmdGo)
	if rootImportPath == "" {
		return
	}
	collector := newPackageImportCollector(cmdGo, rootImportPath)
	for _, mainImpPath := range mainImportPaths {
		if _, err := collector.fetchPackageInfo(mainImpPath); nil != err {
			log.Printf("failed on fetch main package info %v: %v", mainImpPath, err)
			break
		}
	}
	pkgRefs, maxDepth := collector.withImportOrder()
	for _, pkgRef := range pkgRefs {
		log.Printf("%d: %s", maxDepth-pkgRef.DepthValue, pkgRef.ImportPath)
		for _, impByRef := range pkgRef.BeImports {
			log.Printf("  - by: %s", impByRef.ImportPath)
		}
		for _, impsRef := range pkgRef.ImportsOther {
			log.Printf("  - imps: %s", impsRef.ImportPath)
		}
	}
	return
}
