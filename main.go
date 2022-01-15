package main

import (
	"flag"
	"fmt"
	"os"
)

var Usage = func() {
	fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

const (
	TextColor = "\033[38;5;%dm%s\033[39;49m"
	DANGER    = 196
	WARNING   = 208
	SUCCESS   = 2
)

func main() {

	var size float64
	var fileName string

	flag.StringVar(&fileName, "f", "", "File to truncate")
	flag.Float64Var(&size, "s", -1, "Truncate size in megabytes, 0 to empty file")
	flag.Parse()

	if size < 0 {
		fmt.Printf(TextColor, DANGER, "Size is required\n")
		Usage()
		os.Exit(1)
	}

	if len(fileName) == 0 {
		fmt.Printf(TextColor, DANGER, "File is required\n")
		Usage()
		os.Exit(1)
	}

	bytes := (int64)(size * 1024.0 * 1024.0)

	if fi, err := os.Stat(fileName); err != nil || fi.IsDir() {
		fmt.Printf(TextColor, WARNING, "File not found or is a directory.\n")
		os.Exit(1)
	} else {
		bytes = min(bytes, fi.Size())
	}

	if err := os.Truncate(fileName, bytes); err != nil {
		fmt.Printf(TextColor, 196, "Error truncating file: "+err.Error()+"\n")
		os.Exit(1)
	}

	fmt.Printf(TextColor, SUCCESS, "File truncated successfully!\n")

	os.Exit(0)

}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
