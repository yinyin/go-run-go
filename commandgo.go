package rungo

import (
	"errors"
)

// DefaultGoExecutablePath is the default path of `go` executable.
const DefaultGoExecutablePath = "go"

// ErrUnexpectCommandOutput indicate the output of `go` command not in expected form.
// The content parsing is failed.
var ErrUnexpectCommandOutput = errors.New("output of `go` command not in expected form")

// CommandGo wraps command invocation and result parsing of `go` command.
type CommandGo struct {
	ExecutablePath string
}

func (c *CommandGo) exePath() string {
	if "" == c.ExecutablePath {
		c.ExecutablePath = DefaultGoExecutablePath
	}
	return c.ExecutablePath
}
