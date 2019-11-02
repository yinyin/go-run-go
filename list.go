package rungo

// ListPackage get package information with `go list -json`.
func (c *CommandGo) ListPackage(pkgImportPath ...string) (result []*Package, err error) {
	cmdArgs := []string{"list", "-json"}
	cmdArgs = append(cmdArgs, pkgImportPath...)
	jsonResult, err := c.runWithJSONResult(cmdArgs)
	if nil != err {
		return
	}
	defer jsonResult.Wait()
	ok := true
	for ok {
		var aux Package
		if ok, err = jsonResult.Decode(&aux); ok {
			result = append(result, &aux)
		} else if nil != err {
			return
		}
	}
	return
}

// ListModule get module information with `go list -json -m`.
func (c *CommandGo) ListModule(pkgImportPath ...string) (result []*Module, err error) {
	cmdArgs := []string{"list", "-json", "-m"}
	cmdArgs = append(cmdArgs, pkgImportPath...)
	jsonResult, err := c.runWithJSONResult(cmdArgs)
	if nil != err {
		return
	}
	defer jsonResult.Wait()
	ok := true
	for ok {
		var aux Module
		if ok, err = jsonResult.Decode(&aux); ok {
			result = append(result, &aux)
		} else if nil != err {
			return
		}
	}
	return
}
