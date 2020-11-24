package main

import (
	"github.com/toteki/wiz"
	"log"
	"regexp"
	"strings"
)

func main() {

	args := wiz.Args()

	wiz.Green("Reading file " + args[0])
	b, err := wiz.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	s0 := []byte("<script>\n")
	s1 := []byte("\n</script>")

	cli := wiz.NewClient(nil, 10)

	wiz.Green("Fetching file godocs.js")
	godocs, err := cli.Get("https://raw.githubusercontent.com/toteki/doc/main/godocs.js")
	//godocs, err := wiz.ReadFile("godocs.js")
	if err != nil {
		log.Fatal(err)
	}
	godocs = append(s0, godocs...)
	godocs = append(godocs, s1...)

	wiz.Green("Fetching file jquery.js")
	jquery, err := cli.Get("https://raw.githubusercontent.com/toteki/doc/main/jquery.js")
	if err != nil {
		log.Fatal(err)
	}
	jquery = append(s0, jquery...)
	jquery = append(jquery, s1...)

	c0 := []byte("<style>\n")
	c1 := []byte("\n</style>")

	wiz.Green("Fetching file style.css")
	style, err := cli.Get("https://raw.githubusercontent.com/toteki/doc/main/style.css")
	if err != nil {
		log.Fatal(err)
	}
	style = append(c0, style...)
	style = append(style, c1...)

	//
	//
	//

	line := regexp.MustCompile(".*")

	lines := line.FindAll(b, -1)

	lines = purge(lines, "<div id=\"footer\">", "</div>")

	//
	//
	//

	lines = replace(lines, "using GOPATH mode", []byte{})
	lines = replace(lines, "<script src=\"/lib/godoc/jquery.js\" defer></script>", jquery)
	lines = replace(lines, "<script src=\"/lib/godoc/godocs.js\" defer></script>", godocs)
	lines = replace(lines, "<link type=\"text/css\" rel=\"stylesheet\" href=\"/lib/godoc/style.css\">", style)

	//
	//
	//

	out := []byte{}
	for _, l := range lines {
		out = append(out, l...)
		out = append(out, []byte("\n")...)
	}
	wiz.Green("Finished turning into single file")

	//Next, should transform
	//<a href="/src/github.com/toteki/wiz/ASCII.go?s=653:688#L14">ASCII</a>
	//to
	//ASCII
	//following the rule that, if the href starts with /, eliminate

	//The regex that will match them
	refspotter := regexp.MustCompile("<a href=\"/[^\"]*\">")
	out = refspotter.ReplaceAllLiteral(out, []byte("<a doomed href>"))
	wiz.Red("Spotting all broken hyperlinks")

	//The regex that will kill them
	refkiller := regexp.MustCompile("<a doomed href>[^<]*</a>")
	out = refkiller.ReplaceAllFunc(out, striplink)
	wiz.Red("Killing all broker hyperlinks")

	err = wiz.WriteFile(args[1], out)
	if err != nil {
		log.Fatal(err)
	}
	wiz.Green("Wrote file " + args[1])
}

func striplink(in []byte) []byte {
	s := string(in)
	s = strings.TrimPrefix(s, "<a doomed href>")
	s = strings.TrimSuffix(s, "</a>")
	return []byte(s)
}

func replace(lines [][]byte, line string, file []byte) [][]byte {
	//Replaces a given line in a [][]byte line-separated file, with not
	//  one line but an entire byte stream. In this case a file.
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

func purge(lines [][]byte, line1, line2 string) [][]byte {
	//Purges all lines in a [][]byte line-separated file, between two specified
	//  lines of text. (Purge eliminates the specified lines themselves too)
	active := false
	out := [][]byte{}
	for _, l := range lines {
		if !active {
			//If not currently purging between two lines
			if string(l) == line1 {
				active = true //Start purging
			}
		}
		if active {
			//If currently purging between two lines
			if string(l) == line2 {
				active = false //Stop purging
			}
		}
		//If not purging AFTER checks
		if !active {
			out = append(out, l)
		}
	}
	return out
}
