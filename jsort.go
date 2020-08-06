package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	var pretty bool
	flag.BoolVar(&pretty, "pretty", false, "Pretty-print JSON output")
	flag.Parse()

	var v interface{}
	if err := json.NewDecoder(os.Stdin).Decode(&v); err != nil {
		panic(fmt.Sprintf("Could not decode input JSON: %v", err))
	}

	enc := json.NewEncoder(os.Stdout)
	if pretty {
		enc.SetIndent("", "  ")
	}

	if err := enc.Encode(v); err != nil {
		panic(fmt.Sprintf("Could not encode sorted JSON: %v", err))
	}
}
