# doc

The go ecosystem has really good documentation tools (notably godoc and pkg.go.dev), which generate html documentation for go packages, but those don't quite do well generating standalone documentation files from private git repositories.

Specifically, pkg.go.dev requires your code be in a public repository so the tool can read it; and godoc works locally but will generate interdependent html files with links to localhost:6060 that will break once you stop the tool, and which also require internet access to view properly as they use external css and javascript.

This tool uses godoc to generate html documentation for a given package in your $GOPATH (as long as you personally have access to it with 'go get'), embeds the remote css and javascript directly into the html, and culls all hyperlinks that would break once internet access it lost or the godoc tool is turned off.

The process works on private repositories.

The output is a standalone html file, which can be opened normally without internet access or an actively running godoc server.

## Requirements:
- godoc command installed (not go doc, but godoc)
- go installed
- go get -u github.com/toteki/wiz

## How to do it:

- Clone this repository
- Ensure the module you want to document is in your $GOPATH (e.g. using go get)
- Run the doc.sh script
- Open the resulting HTML doc

```
# For example, to generate a document github.com/pkg/errors
git clone https://github.com/toteki/doc
go get github.com/pkg/errors

# doc.sh will prompt for the package name, i.e. github.com/pkg/errors
cd doc
sh doc.sh

# The doc's name is the final term (i.e. after last slash) of package name
open output/errors.html
```

## What doc.sh does:
- Prompts user to enter a package to document, like github.com/toteki/wiz
- Runs godoc to generate html documentation of the package (from $GOPATH)
- Runs doc.go, which fixes the godoc html output (style and broken links)
- Cleans up, leaving the resulting document in the output folder


## TODO
- Create standalone execution mode, which can be run either as a command line tool, or as a wget && sh oneshot script, that works within go modules rather than requiring GOPATH.
- Make an option/flag on the doc.go script, which instead of embedding the js/css, simply points to the versions hosted in toteki/doc so as to reduce html output size.
