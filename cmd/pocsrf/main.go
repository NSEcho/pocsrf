package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lateralusd/pocsrf"
)

func main() {
	input := flag.String("i", "req", "input filename")
	out := flag.String("o", "out_csrf.html", "output filename")
	useHTTPS := flag.Bool("s", false, "use http or https (default http)")
	flag.Parse()

	var schema string
	if *useHTTPS {
		schema = "https://"
	} else {
		schema = "http://"
	}

	f, err := os.Open(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pocsrf: error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	poc, err := pocsrf.NewPOC(f, schema)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pocsrf: error creating poc: %v\n", err)
		os.Exit(1)
	}

	output, err := os.Create(*out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pocsrf: error creating \"%s\": %v\n", *out, err)
		os.Exit(1)
	}
	defer output.Close()

	if err := poc.Write(output); err != nil {
		fmt.Fprintf(os.Stderr, "pocsrf: error writing to \"%s\": %v\n", *out, err)
		os.Exit(1)
	}

	fmt.Println("DONE")
}
