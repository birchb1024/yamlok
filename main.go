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

func makeBooleanFlag(flagVar *bool, swich string, desc string) {
	flag.BoolVar(flagVar, swich, false, desc)
	flag.BoolVar(flagVar, string(swich[0]), false, desc)
}


func helpText() {
usage := `

yamlok takes a list of YAML files as arguments. It parses each file in turn. If an error is found,
processing stops and details are printed on stderr. If all the files are ok the process status is zero
otherwise non zero. 

If no input files are given yamlok reads YAML from the standard input. 

If the --echo option is given, the YAML is also regenerated and sent to stdout. 
`
			fmt.Fprintln(os.Stderr, "Simple program to validate YAML files.\n\nusage:\n\n   yamlok [-h|--help] [-e|--echo] [File...]\n")
			flag.PrintDefaults()
			fmt.Fprintln(os.Stderr, usage)
	
}
func main() {


	var echo, help bool

	makeBooleanFlag(&echo, "echo", "Output the parsed YAML to stdout.")
	makeBooleanFlag(&help, "help", "Print helpful text.")
	 
	flag.Parse()

	if help {
		helpText()
		return
	}

	err := yaml_check(flag.Args(), echo, os.Stdout)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
