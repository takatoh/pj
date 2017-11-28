package main

import (
	"fmt"
	"encoding/json"
	"bytes"
	"os"
	"io/ioutil"
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
	infile := os.Args[1]
	src, err := ioutil.ReadFile(infile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pretty := prettyJson(src)
	fmt.Println(pretty)
}
