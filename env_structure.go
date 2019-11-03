package rungo

// EnvInfo represent the result of `go env` command.
// JSON key came from: https://golang.org/src/cmd/go/internal/envcmd/env.go
type EnvInfo struct {
	Arch       string `json:"GOARCH"`
	Bin        string `json:"GOBIN"`
	Cache      string `json:"GOCACHE"`
	EnvFile    string `json:"GOENV"`
	ExeSuffix  string `json:"GOEXE"`
	Flags      string `json:"GOFLAGS"`
	HostArch   string `json:"GOHOSTARCH"`
	HostOS     string `json:"GOHOSTOS"`
	GoNoProxy  string `json:"GONOPROXY"`
	GoNoSumDB  string `json:"GONOSUMDB"`
	OS         string `json:"GOOS"`
	GoPath     string `json:"GOPATH"`
	GoPrivate  string `json:"GOPRIVATE"`
	GoProxy    string `json:"GOPROXY"`
	GoRoot     string `json:"GOROOT"`
	GoSumDB    string `json:"GOSUMDB"`
	TmpDir     string `json:"GOTMPDIR"`
	ToolDir    string `json:"GOTOOLDIR"`
	GCCGO      string `json:"GCCGO"`
	AR         string `json:"AR"`
	CC         string `json:"CC"`
	CXX        string `json:"CXX"`
	CGOENABLED string `json:"CGO_ENABLED"`
}
