package main

import (
	"fmt"
	"encoding/json"
	"bytes"
	"os"
	"io/ioutil"
	"flag"
)

const (
	progVersion = "v0.1.0"
)

func prettyJson(src []byte) string {
	buf := make([]byte, 0)
	dst := bytes.NewBuffer(buf)
	err := json.Indent(dst, src, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return dst.String()	
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
`Usage:
  %s [options] <file.json>
Options:
`, os.Args[0])
		flag.PrintDefaults()
	}
	opt_version := flag.Bool("version", false, "Show version.")
	flag.Parse()

	if *opt_version {
		fmt.Println(progVersion)
		os.Exit(0)
	}

	infile := flag.Args()[0]
	src, err := ioutil.ReadFile(infile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pretty := prettyJson(src)
	fmt.Println(pretty)
}
