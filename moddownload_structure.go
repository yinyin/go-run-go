package rungo

// The following structure came from (LICENSE attached at tail of code):
// https://golang.org/src/cmd/go/internal/modcmd/download.go

// ModuleDownloadResult is the structure of module download description
// structure in `go mod download -json` result.
type ModuleDownloadResult struct {
	Path     string `json:",omitempty"` // module path
	Version  string `json:",omitempty"` // module version
	Error    string `json:",omitempty"` // error loading module
	Info     string `json:",omitempty"` // absolute path to cached .info file
	GoMod    string `json:",omitempty"` // absolute path to cached .mod file
	Zip      string `json:",omitempty"` // absolute path to cached .zip file
	Dir      string `json:",omitempty"` // absolute path to cached source root directory
	Sum      string `json:",omitempty"` // checksum for path, version (as in go.sum)
	GoModSum string `json:",omitempty"` // checksum for go.mod (as in go.sum)
	Latest   bool   // would @latest resolve to this version?
}

// License information:
// https://golang.org/LICENSE
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
