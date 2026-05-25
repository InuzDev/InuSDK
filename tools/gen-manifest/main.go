package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: gen-manifest <sdk> <version>")
		fmt.Println("Example: gen-manifest java 21.0.2+13")
		os.Exit(1)
	}

	sdk := os.Args[1]
	version := os.Args[2]

	switch sdk {
	case "java":
		_manifest, err := generateJava(version)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		output, err := json.MarshalIndent(_manifest, "", " ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshalling JSON: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
	default:
		fmt.Fprintf(os.Stderr, "Unsupported SDK: %s\n", sdk)
		os.Exit(1)
	}
}
