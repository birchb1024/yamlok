package main

import (
	"os"
	"strings"
	"testing"
)

func emitNothing(lines string) {}

func Test_check_file_echo(t *testing.T) {

	var result string
	emit := func(lines string) {
		result += lines
	}
	err := check_file(strings.NewReader("{ a: 1, b: 2 }"), emit)
	expected := "---\na: 1\nb: 2\n...\n"
	if result != expected {
		t.Errorf("expected result: %#v", expected)
		t.Errorf("actual result: %#v", result)
	}
	if err != nil {
		t.Errorf("unexpected error:" + err.Error())
	}
}

func Test_check_file_no_echo(t *testing.T) {

	err := check_file(strings.NewReader("{ a: 1, b: 2 }"), emitNothing)
	if err != nil {
		t.Errorf("unexpected error:" + err.Error())
	}
}

func Test_check_bad_file_no_echo(t *testing.T) {

	err := check_file(strings.NewReader("{ a: 1 : b: 2 }"), emitNothing)
	if err == nil {
		t.Errorf("was expecting an error")
	}
	if !strings.Contains(err.Error(), "expected ',' or '}'") {
		t.Errorf("did not find %#v", err.Error())
	}
}

func Test_check_good_files_no_echo(t *testing.T) {
	var output strings.Builder
	err := yaml_check([]string{"test/testdata/good1.yaml", "test/testdata/good2.yaml"}, false, &output)
	if err != nil {
		t.Errorf("was not expecting an error %#v", err.Error())
	}
}

func Test_check_good_files_echo(t *testing.T) {
	var output strings.Builder
	err := yaml_check([]string{"test/testdata/good1.yaml", "test/testdata/good2.yaml"}, true, &output)
	if err != nil {
		t.Errorf("was not expecting an error %#v", err.Error())
	}
	expected := "# test/testdata/good1.yaml\n---\n- 2\n- 3\n- four: 5\n...\n# test/testdata/good2.yaml\n---\n- 2\n- 3\n- four: 5\n...\n"
	if output.String() != expected {
		t.Errorf("expected result: %#v", expected)
		t.Errorf("actual result: %#v", output.String())
	}
}

func Test_check_bad_files_echo(t *testing.T) {
	var output strings.Builder
	err := yaml_check([]string{"test/testdata/good1.yaml", "test/testdata/bad1.yaml"}, true, &output)
	if err == nil {
		t.Errorf("was expecting an error %#v", err.Error())
	}
	if !strings.Contains(err.Error(), "expected '-' indicator") {
		t.Errorf("did not find %#v", err.Error())
	}
}

func Test_check_bad_filenames_echo(t *testing.T) {
	var output strings.Builder
	err := yaml_check([]string{"test/testdata/good1.yaml", "test/testdata/missing.yaml"}, true, &output)
	if err == nil {
		t.Errorf("was expecting an error %#v", err.Error())
	}
	if !strings.Contains(err.Error(), "no such") {
		t.Errorf("did not find %#v", err.Error())
	}
}

func Test_check_good_stdin_echo(t *testing.T) {
	file, err := os.Open("test/testdata/good1.yaml")
	if err != nil {
		t.Errorf(err.Error())
	}
	os.Stdin = file // Mock Stdin

	var output strings.Builder
	err = yaml_check([]string{}, true, &output)
	if err != nil {
		t.Errorf("was not expecting an error %#v", err.Error())
	}
	expected := "---\n- 2\n- 3\n- four: 5\n...\n"
	if output.String() != expected {
		t.Errorf("expected result: %#v", expected)
		t.Errorf("actual result: %#v", output.String())
	}
}

func Test_bad_stdin_echo(t *testing.T) {
	file, err := os.Open("test/testdata/bad1.yaml")
	if err != nil {
		t.Errorf(err.Error())
	}
	os.Stdin = file // Mock Stdin

	var output strings.Builder
	err = yaml_check([]string{}, true, &output)
	if err == nil {
		t.Errorf("was expecting an error")
	}
	if !strings.Contains(err.Error(), "expected '-'") {
		t.Errorf("did not find %#v", err.Error())
	}
}

func Test_raise_coverage_hehe(t *testing.T) {
	var output1, output2 strings.Builder
	helpText(&output1, false)
	if output1.String() != "" {
		t.Errorf("got help when not expected %#v", output1.String())
	}
	helpText(&output2, true)
	if len(output2.String()) < 40 {
		t.Errorf("got no help when expected %#v", output2.String())
	}
}
