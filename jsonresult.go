package rungo

import (
	"encoding/json"
	"io"
	"os/exec"
)

type jsonResultDecoder struct {
	cmd    *exec.Cmd
	stdout io.ReadCloser
	dec    *json.Decoder
}

func (r *jsonResultDecoder) Decode(v interface{}) (ok bool, err error) {
	if err = r.dec.Decode(v); nil != err {
		if io.EOF == err {
			return false, nil
		}
		return
	}
	return true, nil
}

func (r *jsonResultDecoder) Wait() error {
	return r.cmd.Wait()
}

func (c *CommandGo) runWithJSONResult(cmdArgs []string) (resDecoder *jsonResultDecoder, err error) {
	cmd := exec.Command(c.exePath(), cmdArgs...)
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		return
	}
	if err = cmd.Start(); nil != err {
		return
	}
	dec := json.NewDecoder(stdout)
	resDecoder = &jsonResultDecoder{
		cmd:    cmd,
		stdout: stdout,
		dec:    dec,
	}
	return
}
