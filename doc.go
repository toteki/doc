package main

import (
	"github.com/toteki/wiz"
	"log"
	"regexp"
	"strings"
)

func main() {

	// Only run this file after acquiring a raw godoc output html.
	// One way to do this is by:

	/*
		read -p "Enter package url to document: " PACKAGE
		godoc -url "http://localhost:6060/pkg/$PACKAGE/" > raw_doc.html
	*/

	// This script takes raw godoc output HTML, which contains references to
	// external CSS and JS, as well as links to localhost:6060, where godoc
	// would be running, and proceeds to embed the external files and clean
	// any hyperlinks that wouldn't work without once the godoc server stopped.
	// The CSS and JS are stored in this repository to remove external dependency.
	// They were retrieved November 2020 and will not be updated unless necessary.

	args := wiz.Args() // Retrieve command line arguments (expects 2)

	// The command line argument is the name of the raw html doc to clean.
	// This is doc_raw.html when using as intended.
	wiz.Green("Reading file " + args[0])
	raw_file, err := wiz.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	// HTML script tags
	s0 := []byte("<script>\n")
	s1 := []byte("\n</script>")

	// Read godocs.js and prepend+append script tags
	wiz.Green("Reading file godocs.js")
	godocs, err := wiz.ReadFile("godocs.js")
	if err != nil {
		log.Fatal(err)
	}
	godocs = append(s0, godocs...)
	godocs = append(godocs, s1...)

	// Read jquery.js and prepend+append script tags
	wiz.Green("Reading file jquery.js")
	jquery, err := wiz.ReadFile("jquery.js")
	if err != nil {
		log.Fatal(err)
	}
	jquery = append(s0, jquery...)
	jquery = append(jquery, s1...)

	// HTML style tags
	c0 := []byte("<style>\n")
	c1 := []byte("\n</style>")

	// Read style.css and prepend+append style tags
	wiz.Green("Reading file style.css")
	style, err := wiz.ReadFile("style.css")
	if err != nil {
		log.Fatal(err)
	}
	style = append(c0, style...)
	style = append(style, c1...)

	// Separate file out into lines
	line := regexp.MustCompile(".*")    // A regex which matches anything
	lines := line.FindAll(raw_file, -1) // All matching lines in the file

	// Eliminate the footer of the HTML document
	lines = eliminateRange(lines, "<div id=\"footer\">", "</div>")

	// Eliminate unnecessary godoc message
	lines = replace(lines, "using GOPATH mode", []byte{})
	// Embed local jquery.js file
	lines = replace(lines, "<script src=\"/lib/godoc/jquery.js\" defer></script>", jquery)
	// Embed local godocs.js file
	lines = replace(lines, "<script src=\"/lib/godoc/godocs.js\" defer></script>", godocs)
	// Embed local style.css file
	lines = replace(lines, "<link type=\"text/css\" rel=\"stylesheet\" href=\"/lib/godoc/style.css\">", style)

	// Re-combine the file line by line after manipulation
	out := []byte{}
	for _, l := range lines {
		out = append(out, l...)
		out = append(out, []byte("\n")...)
	}
	wiz.Green("Finished line manipulation")

	// Next, should transform any hyperlinks in the file to plain text.
	//
	// For example,
	// <a href="/src/github.com/toteki/wiz/ASCII.go?s=653:688#L14">ASCII</a>
	// becomes
	// ASCII
	//
	// Only eliminates hrefs that start with / (which leaves intact the ones
	// which only jump the user to a section of the document).

	// Mark these references to outside the document by
	// replacing their first tag with <a doomed href>
	refspotter := regexp.MustCompile("<a href=\"[^#][^\"]*\">")
	out = refspotter.ReplaceAllLiteral(out, []byte("<a doomed href>"))
	wiz.Purple("Spotting all broken hyperlinks")

	// Kill <a doomed href> tags while
	// preserving what's between them.
	refkiller := regexp.MustCompile("<a doomed href>[^<]*</a>")
	out = refkiller.ReplaceAllFunc(out, striplink)
	wiz.Purple("Killing all broken hyperlinks")

	// Determine the filename out the output - we will use everything after the
	// last slash, e.g. "github.com/pkg/errors.html" => "errors.html"

	// Matches zero or more non-slash before end of string
	namer := regexp.MustCompile("[^/]+$")
	name := string(namer.Find([]byte(args[1]))) // use on second command line arg

	// Writing the cleaned output file to output directory
	err = wiz.WriteFile("output/"+name, out)
	if err != nil {
		log.Fatal(err)
	}
	wiz.Green("Wrote file " + name + " to output directory")
}

// Given an input string which begins and ends with <a doomed href> and </a>,
// remove those tags. This is used after those tags are inserted in place
// of hyperlinks this script wishes to remove.
func striplink(in []byte) []byte {
	s := string(in)
	s = strings.TrimPrefix(s, "<a doomed href>")
	s = strings.TrimSuffix(s, "</a>")
	return []byte(s)
}

// Replaces a given line in a [][]byte representing the lines in a file, with
// a given byte slice. (In this case, a whole file.)
func replace(lines [][]byte, line string, file []byte) [][]byte {
	out := [][]byte{}
	for _, l := range lines {
		if string(l) == line {
			wiz.Purple("Replacing line: ", line)
			out = append(out, file) //add the file
		} else {
			out = append(out, l) //add the existing line
		}
	}
	return out
}

// Removes all lines between two specified lines of text in a [][]byte
// representing the lines in a file. Eliminates the specified lines too.
func eliminateRange(lines [][]byte, line1, line2 string) [][]byte {
	active := false
	output := [][]byte{}
	for _, l := range lines {
		if !active {
			//If not currently eliminating between two lines
			if string(l) == line1 {
				active = true //Start eliminating if encountered line1
			}
		}
		//If not eliminating, add lines to output
		if !active {
			output = append(output, l)
		}
		if active {
			//If currently eliminating between two lines
			if string(l) == line2 {
				active = false //Stop eliminating if encountered line2
			}
		}
	}
	return output
}
