package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

func yaml_check(filenames []string, echo bool, output io.Writer) (err error) {

	emit := func(lines string) {
		if echo {
			fmt.Fprint(output, lines)
		}
	}

	if len(filenames) == 0 {
		err = check_file(os.Stdin, emit)
		if err != nil {
			return err
		}
	}
	
	for _, filename := range filenames {
		input, err := os.Open(filename)
		if err != nil {
			return err
		}

		if len(filenames) > 1 {
			emit("# " + filename + "\n")
		}
		err = check_file(input, emit)
		if err != nil {
			return errors.WithMessage(err, filename)
		}
	}
	return nil
}

func check_file(input io.Reader, emit func(string)) (err error) {

	decoder := yaml.NewDecoder(input)
	for {
		var data interface{}
		err := decoder.Decode(&data)
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		yamldata, yerr := yaml.Marshal(data)
		if yerr != nil {
			return err
		}
		emit("---\n")
		emit(string(yamldata))
		emit("...\n")
	}
	return nil
}

func main() {

	var echo bool

	flag.BoolVar(&echo, "echo", false, "Output the parsed YAML to stdout.")
	flag.BoolVar(&echo, "e", false, "Output the parsed YAML to stdout.")

	flag.Parse()

	err := yaml_check(flag.Args(), echo, os.Stdout)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
