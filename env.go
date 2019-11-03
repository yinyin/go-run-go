package rungo

// Env get environment information with `go env -json`.
func (c *CommandGo) Env() (result *EnvInfo, err error) {
	cmdArgs := []string{"env", "-json"}
	var envInfo EnvInfo
	if err = c.runWithJSONDecode(cmdArgs, &envInfo); nil != err {
		return
	}
	result = &envInfo
	return
}
