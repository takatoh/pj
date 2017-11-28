package main

import (
	"fmt"
	"encoding/json"
	"bytes"
	"os"
	"io/ioutil"
	"flag"
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
  %s <file.json>
`, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	infile := flag.Args()[0]
	src, err := ioutil.ReadFile(infile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pretty := prettyJson(src)
	fmt.Println(pretty)
}
