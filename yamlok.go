package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

func main() {

	var echo bool
	flag.BoolVar(&echo, "echo", false, "Output the parsed YAML to stdout.")

	flag.Parse()
	if len(flag.Args()) == 0 {
			log.Fatalf("error: no files to parse given: %v", os.Args)
	}

	for _, filename := range flag.Args() {
		input, err := os.Open(filename)
		if err != nil {
			log.Fatalf("bad filename '%v':  %v", filename, err)
		}
		decoder := yaml.NewDecoder(input)

		if len(flag.Args()) > 1 {
			fmt.Println("#", filename)
		}
		for {
			var data interface{}
			err := decoder.Decode(&data)
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("%v error: %v", filename, err)
			}
			yamldata, yerr := yaml.Marshal(data)
			if yerr != nil {
				log.Fatalf("%v error: %v", filename, err)
			}
			if echo {
				fmt.Println("---")
				fmt.Print(string(yamldata))
				fmt.Println("...")
			}
		}
	}
}
