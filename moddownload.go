package rungo

// ModuleDownload downloads the named modules with `go mod download -json -m`.
func (c *CommandGo) ModuleDownload(modulePathPatterns ...string) (result []*ModuleDownloadResult, err error) {
	cmdArgs := []string{"mod", "download", "-json"}
	cmdArgs = append(cmdArgs, modulePathPatterns...)
	jsonResult, err := c.runWithJSONResults(cmdArgs)
	if nil != err {
		return
	}
	defer jsonResult.Wait()
	ok := true
	for ok {
		var aux ModuleDownloadResult
		if ok, err = jsonResult.Decode(&aux); ok {
			result = append(result, &aux)
		} else if nil != err {
			return
		}
	}
	return
}
