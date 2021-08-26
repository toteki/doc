#!sh

# This script is the version of the tool that would be run from within this
# repository once cloned (as opposed to the standalone script which can be
# run without cloning). See README for detailed intructions.
read -p "Enter package url to document: " PACKAGE

# Create the file raw_doc.html by funneling godoc output into a file
godoc -url "http://localhost:6060/pkg/$PACKAGE/" > raw_doc.html

# Build doc.go (Unix won't care about the .exe)
go build -o doc.exe doc.go

# Run the executable, passing the raw_doc filename and the package name as args
./doc.exe raw_doc.html $PACKAGE.html
# This produces an html file in the output directory

# Remove the raw doc and the executable
rm doc.exe
rm raw_doc.html
