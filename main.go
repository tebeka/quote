package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
)

//go:generate go run gen_quotes.go
//go:generate go fmt quotes.go

func main() {
	var count int
	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "usage: %s [options]\n", name)
		fmt.Fprintln(os.Stderr, "Will print `count` random quotes.\nOptions:")
		flag.PrintDefaults()
	}
	flag.IntVar(&count, "n", 1, "number of results to return")
	flag.Parse()

	if flag.NArg() > 1 {
		fmt.Fprintf(os.Stderr, "error: wrong number of arguments\n")
		os.Exit(1)
	}

	quotes := quoteDB

	if flag.NArg() == 1 {
		query := strings.ToLower(flag.Arg(0))
		quotes = make([]string, 0, len(quoteDB))
		for _, q := range quoteDB {
			text := strings.ToLower(q)
			if strings.Contains(text, query) {
				quotes = append(quotes, q)
			}
		}

		if len(quotes) == 0 {
			fmt.Fprintf(os.Stderr, "error: not quotes matching %q\n", query)
			os.Exit(1)
		}
	}

	rand.Shuffle(len(quotes), func(i, j int) {
		quotes[i], quotes[j] = quotes[j], quotes[i]
	})

	n := min(count, len(quotes))
	for i := range n {
		fmt.Println(quotes[i])
		fmt.Println()
	}
}
