package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	highlight "github.com/d4l3k/go-highlight"
	"github.com/fatih/color"
)

var (
	html = flag.Bool("html", false, "output html")
	term = flag.Bool("term", false, "output terminal colors")
	lang = flag.String("lang", "go", "specify the language")
)

// http://golang-examples.tumblr.com/post/97848675424/read-from-stdin-or-file
func openStdinOrFile() io.Reader {
	var err error
	r := os.Stdin
	args := flag.Args()
	if len(args) > 0 {
		r, err = os.Open(args[0])
		if err != nil {
			panic(err)
		}
	}
	return r
}

func main() {
	flag.Parse()

	if *html && *term {
		log.Fatal("can't specify -html and -term")
	}

	in := openStdinOrFile()
	bytes, err := ioutil.ReadAll(in)
	if err != nil {
		log.Fatal(err)
	}

	highlightFunc := highlight.Term
	if *html {
		log.Println("html!")
		highlightFunc = highlight.HTML
	}

	highlighted, err := highlightFunc(*lang, string(bytes))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(color.Output, highlighted)
}
