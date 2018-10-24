package main

import (
	"fmt"
	"io"
	"text/template"
)

func mainUsage(f io.Writer) {
	t := template.Must(template.New("").Parse(mainHelpTemplate))
	if err := t.Execute(f, nil); err != nil {
		fmt.Fprintf(f, "cannot write usage output: %v", err)
	}
}

var mainHelpTemplate = `
The gobin command installs/runs main packages.

Usage:
	gobin [-m] [-r|-p] [-n|-g] packages [run arguments...]

The packages argument to gobin is similar to that of the go tool (in module
mode) with the additional constraint that the list of packages must be main
packages that are part of Go modules.

By default, gobin is said to operate in global mode. If the -m flag is provided
then it is said to operate in main-module mode, where the path to the main
module's go.mod is given by go env GOMOD.

The version "latest" matches the latest available tagged version. If no version
is specified, gobin behaves differently in global and main-module modes. In
global mode, gobin attempts to resolve the latest version available in the
module download cache. In main-module module, gobin attempts to resolve the
current version via the main module's go.mod. If this resolution fails in
either mode, "latest" is assumed and gobin resolves via the network.

In global mode, gobin installs the main packages to the directories
gobin/$module@$version/$main_pkg under your user cache directory. See the
documentation for os.UserCacheDir for OS-specific details on how to configure
its location.

In main-module mode, gobin installs the main packages to the directories
.gobincache/$module@$version/$main_pkg under the directory containing the main
module's go.mod.

By default, gobin installs the main packages to $GOBIN (or $GOPATH/bin if GOBIN
is not set, which defaults to $HOME/go/bin if GOPATH is not set).

The -r flag takes exactly one main package argument and runs that package.  It
is similar therefore to go run. Any arguments after the single main package
will be passed to the main package as command line arguments.

The -p flag prints the cache install path for each of the provided packages
once versions have been resolved.

The -r and -p flags are mutually exclusive.

The -g flag forces gobin to perform a network fetch for the provided main
packages.

The -n flag prevents network access and instead uses the GOPATH module download
cache where required.

The -g and -n flags are mutually exclusive.

The -m flag causes gobin to use the main module (the module containing the
directory where the gobin command is run). The main module is given by go env
GOMOD. Without this flag gobin effectively runs as a "global" tool.

It is an error for a non-main package, a package pattern or a main package not
part of a module to be provided as a package argument.

To install a main package not part of a go module, use:

	 GO111MODULE=off go get main_pkg

`[1:]