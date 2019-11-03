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

func (c *CommandGo) runWithJSONResults(cmdArgs []string) (resDecoder *jsonResultDecoder, err error) {
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

func (c *CommandGo) runWithJSONDecode(cmdArgs []string, v interface{}) (err error) {
	cmd := exec.Command(c.exePath(), cmdArgs...)
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		return
	}
	if err = cmd.Start(); nil != err {
		return
	}
	dec := json.NewDecoder(stdout)
	if err = dec.Decode(v); nil != err {
		return
	}
	devnull := make([]byte, 128)
	for {
		if _, err = stdout.Read(devnull); nil != err {
			if err == io.EOF {
				err = nil
			}
			break
		}
	}
	cmd.Wait()
	return
}
