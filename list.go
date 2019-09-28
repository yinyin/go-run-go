package rungo

import (
	"encoding/json"
	"io"
	"os/exec"
)

// ListPackage get package information with `go list -json`.
func (c *CommandGo) ListPackage(pkgImportPath ...string) (result []*Package, err error) {
	cmdArgs := []string{"list", "-json"}
	cmdArgs = append(cmdArgs, pkgImportPath...)
	cmd := exec.Command(c.exePath(), cmdArgs...)
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		return
	}
	if err = cmd.Start(); nil != err {
		return
	}
	dec := json.NewDecoder(stdout)
	for {
		var aux Package
		if err = dec.Decode(&aux); nil != err {
			if io.EOF == err {
				break
			}
			return
		}
		result = append(result, &aux)
	}
	err = cmd.Wait()
	return
}
