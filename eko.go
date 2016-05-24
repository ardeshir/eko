package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)
var out io.Writer = os.Stdout // modified

func main() {
	flag.Parse()
	if err := eko(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "eko: %v\n", err)
		os.Exit(1)
	}
}

func eko(newline bool, sep string, args []string) error {
	fmt.Fprintf(out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}
