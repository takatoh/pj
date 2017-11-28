package main

import (
	"fmt"
	"encoding/json"
	"bytes"
	"os"
	"io/ioutil"
)

func main() {
	infile := os.Args[1]
	src, err := ioutil.ReadFile(infile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	buf := make([]byte, 0)
	dst := bytes.NewBuffer(buf)
	err = json.Indent(dst, src, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(dst.String())
}
