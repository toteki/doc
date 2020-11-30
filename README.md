# doc

Created to be able to generate godoc html pages, except as a flat file without having to serve.

## Requirements:
- godoc installed (not go doc, but godoc) and usable
- go itself (to build the doc.go script)
- go get -u github.com/toteki/wiz

## How to do it:
```
wget -q -O doc https://raw.githubusercontent.com/toteki/doc/main/doc && sh doc
```

## What it does:
- Prompts user to enter a package to document, like github.com/toteki/wiz
- Uses 'go get -u' on that package to bring it up to date
- Runs godoc to generate html documentation of the same package (as retrieved from $GOPATH)
- Builds the doc.go script here in this repository
- Runs the compiled doc.go tool, which fixes the godoc html output (style and broken links)
- Displays the resulting docfinal.html

## Annoyances:
- Since this uses 'go get -u' to fetch the package, you can't use the tool on local changes, or branch changes. The clunky but effective solution to this might just be to push changes to master (or the target branch of 'go get'), THEN run this and push documentation.

- The one-line command might not work on windows (since it ends in a 'sh' command)

- If your system asks for credentials when you use 'go get', then this will occur during the running of the script, probably.
