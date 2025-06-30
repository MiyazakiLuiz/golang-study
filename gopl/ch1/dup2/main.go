// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type countFiles struct {
	c int
	f []string
}

func main() {
	counts := make(map[string]countFiles)
	files := os.Args[1:]
	if len(files) == 0 {
		//countLines(os.Stdin, counts)
		panic("Must inform at least one parameter")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.c > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.c, line, n.f)
		}
	}
}

func countLines(f *os.File, counts map[string]countFiles, file string) {
input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if val, ok := counts[text]; ok {
			val.c++
			val.f = append(val.f, file)
			counts[text] = val
		} else {
			counts[text] = countFiles{
				c: 1,
				f: []string{file},
			}
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-