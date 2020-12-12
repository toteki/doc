# doc

The go ecosystem has really good documentation tools (notably godoc and pkg.go.dev), which generate html documentation for go packages, but those don't quite do well generating standalone documentation files from private git repositories.

Specifically, pkg.go.dev requires your code be in a public repository so the tool can read it; and godoc works locally but will generate interdependent html files with links to localhost:6060 that will break once you stop the tool, and which also require internet access to view properly as they use external css and javascript.

This tool uses godoc to generate html documentation for a given package in your $GOPATH (as long as you personally have access to it with 'go get'), embeds the remote css and javascript directly into the html, and culls all hyperlinks that would break once internet access it lost or the godoc tool is turned off.

The process works on private repositories.

The output is a standalone html file, which can be opened normally without internet access or an actively running godoc server.

## Requirements:
- godoc installed (not go doc, but godoc) and usable
- go itself (to build the doc.go script)
- go get -u github.com/toteki/wiz

## How to do it:
The following grabs the 'doc' shell script from this repository, and runs it with 'sh doc'.
```
wget -q -O doc https://raw.githubusercontent.com/toteki/doc/main/doc && sh doc
```
Then enter a package name when prompted.

## What it does:
- Prompts user to enter a package to document, like github.com/toteki/wiz
- Uses 'go get -u' on that package to bring it up to date (targets master/main branch)
- Runs godoc to generate html documentation of the same package (as retrieved from $GOPATH)
- Builds the doc.go script here in this repository
- Runs the compiled doc.go tool, which fixes the godoc html output (style and broken links)
- Displays the resulting docfinal.html

## Annoyances:
- Since this uses 'go get -u' to fetch the package, you can't use the tool on local changes, or branch changes. The clunky but effective solution to this might just be to push changes to master (or the target branch of 'go get'), THEN run this and push documentation. Otherwise, run the script after commenting out the go get line (assuming your local changes are already in $GOPATH).

- The one-line command might not work on windows command prompt (since it ends in a 'sh' command). Try powershell.

- If your system asks for credentials when you use 'go get', then this will occur during the running of the script, probably.
