//
// pairtree.go implements a command line pairtree utility.
//
// Author R. S. Doiel, <rsdoiel@library.caltech.edu>
//
// Copyright (c) 2021, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/caltechlibrary/pairtree"
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool
)

func usage(appName string, exitCode int) {
	fp := os.Stderr
	if exitCode == 0 {
		fp = os.Stdout
	}
	fmt.Fprintf(fp, `
USAGE %s [OPTIONS] encode|decode STRING

%s will encode or decode a string to/from a pairtree path

OPTIONS

    -h, -help     display help
	-version      display version
	-license      display license
	-s SEP        set separator to SEP

EXAMPLE

Encode key 12345

    %s encode 12345

Decode path 12/34/5

    %s decode 12/34/5

`, appName, appName, appName, appName)
	os.Exit(exitCode)
}

func main() {
	appName := path.Base(os.Args[0])
	sep := fmt.Sprintf("%c", os.PathSeparator)
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showVersion, "version", false, "display version info")
	flag.BoolVar(&showLicense, "license", false, "display license info")
	flag.StringVar(&sep, "s", sep, "set path separator")
	flag.Parse()
	if showHelp {
		usage(appName, 0)
	}
	if showVersion {
		fmt.Fprintf(os.Stdout, "%s %s\n", appName, pairtree.Version)
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(os.Stdout, "%s, %s\n\n", appName, pairtree.Version)
		fmt.Fprintf(os.Stdout, "%s\n", pairtree.License)
		os.Exit(0)
	}
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Wrong number of parameters\n")
		usage(appName, 1)
	}
	if sep[0] != os.PathSeparator {
		pairtree.Set(rune(sep[0]))
	}
	switch strings.ToLower(args[0]) {
	case "decode":
		fmt.Fprintf(os.Stdout, "%s", pairtree.Decode(args[1]))
	case "encode":
		fmt.Fprintf(os.Stdout, "%s", pairtree.Encode(args[1]))
	default:
		fmt.Fprintf(os.Stderr, "Did not understand how to %q\n", args[0])
		usage(appName, 1)
	}
}
