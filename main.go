package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
)

//go:generate go run gen_quotes.go
//go:generate go fmt quotes.go

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func printJSON(quotes []string) {
	var jsonQuotes []Quote
	for _, q := range quotes {
		parts := strings.Split(q, "\n")
		text := strings.Join(parts[:len(parts)-1], " ")
		author := strings.TrimPrefix(parts[len(parts)-1], "    - ")
		jsonQuotes = append(jsonQuotes, Quote{Text: text, Author: author})
	}
	jsonData, err := json.MarshalIndent(jsonQuotes, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}

func main() {
	var count int
	var jsonOutput bool
	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "usage: %s [options]\n", name)
		fmt.Fprintln(os.Stderr, "Will print `count` random quotes.\nOptions:")
		flag.PrintDefaults()
	}
	flag.IntVar(&count, "n", 1, "number of results to return")
	flag.BoolVar(&jsonOutput, "json", false, "output in JSON format")
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
	selectedQuotes := quotes[:n]

	if jsonOutput {
		printJSON(selectedQuotes)
	} else {
		for _, q := range selectedQuotes {
			fmt.Println(q)
			fmt.Println()
		}
	}
}
