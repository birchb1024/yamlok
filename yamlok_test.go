package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_check_file(t *testing.T) {

	// func check_file(input io.Reader, emit func(string)) (err error) {
	var result string
	emit := func(lines string) { 
		result += lines 
	}
	err := check_file(strings.NewReader("{ a: 1, b: 2 }"), emit)
    fmt.Println(result)
	if err != nil {
		t.Errorf("oops")
	}
}
