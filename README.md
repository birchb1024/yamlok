# yamlok

Simple program to validate YAML files.

usage:

```
   yamlok [-h|--help] [-e|--echo] [File...]

  -e	Output the parsed YAML to stdout.
  -echo
    	Output the parsed YAML to stdout.
  -h	Print helpful text.
  -help
    	Print helpful text.
```

yamlok takes a list of YAML files as arguments. It parses each file in turn. If an error is found,
processing stops and details are printed on stderr. If all the files are ok the process status is zero
otherwise non zero. 

If no input files are given yamlok reads YAML from the standard input. 

If the --echo option is given, the YAML is also regenerated and sent to stdout. 

