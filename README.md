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
- Builds to go.doc tool here
- Runs the go.doc tool, which fixes the godoc html output (style and broken links)
- Displays the resulting docfinal.html

## Annoyances:
- Since this uses go get to fetch the package, you can't use the tool on local changes, or branch changes. The clunky but effective solution to this might just be to push changes to master (or the target branch of 'go get'), THEN run this and push documentation.
