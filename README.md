# godoc2ghmd

godoc2ghmd is a main package.
godoc2ghmd (godoc-to-GitHub-MarkDown) generates package documentation in
GitHub flavoured Markdown.

Install:

	go get <a href="https://github.com/DevotedHealth/godoc2ghmd">https://github.com/DevotedHealth/godoc2ghmd</a>

This will install a binary into your go path.

Usage:

	godoc2ghmd [options] [full package name]

For full and up-to-date Usage information run:

	$ godoc2ghmd -help

### Notes
While the output of godoc2ghmd can simply be piped into a file, e.g.:

	$ godoc2ghmd github.com/GandalfUK/godoc2ghmd > README.md

it can also be used with the -file option:

	$ godoc2ghmd -file=README.md github.com/GandalfUK/godoc2ghmd

This invocation is particularly useful when using `go generate` to automate
the creation of package documentation before sumbmitting code. For
example, this directive:

go:generate godoc2ghmd -file=README.md github.com/DevotedHealth/godoc2ghmd

in the  `documentation.go` file within this repositoy created this very
`README.md` file by running:

	$ go generate

The same command also (re)generates all the documentation in the examples
folder via other similar directives.

### History
This is a fork of <a href="https://github.com/GandalfUK/godoc2ghmd">https://github.com/GandalfUK/godoc2ghmd</a>, which itself is
a fork of <a href="https://github.com/davecheney/godoc2md">https://github.com/davecheney/godoc2md</a> and incorporating changes
from <a href="https://github.com/wdamron/godoc2gh">https://github.com/wdamron/godoc2gh</a>.

## <a name="Subdirectories">Subdirectories</a>

* [examples](./examples)
    * [build](./examples/build)
    * [doc](./examples/doc)
    * [http](./examples/http)
    * [ioutil](./examples/ioutil)