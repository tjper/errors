package errors_test

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
	"testing"

	"github.com/tjper/errors"
	testutil "github.com/tjper/testing"
)

var golden = flag.Bool("golden", false, "instructs testing program to overwrite .golden files with test results")

func TestNewErrors(t *testing.T) {
	tests := []struct {
		Name       string
		Processors []errors.Processor
	}{
		{"basic", []errors.Processor{printErr}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			e := errors.NewErrors(test.Processors...)
			if e == nil {
				t.Fatalf("failed %s", test.Name)
			}

			check(t, e.Debug())
		})
	}
}

func TestWith(t *testing.T) {
	tests := []struct {
		Name             string
		InitialProcessor errors.Processor
		Processors       []errors.Processor
	}{
		{"basic", helloWorld, []errors.Processor{printErr}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			e := errors.NewErrors(test.InitialProcessor)
			e.With(test.Processors...)

			check(t, e.Debug())
		})
	}
}

func TestUse(t *testing.T) {
	tests := []struct {
		Name             string
		InitialProcessor errors.Processor
		Processors       []errors.Processor
	}{
		{"basic", helloWorld, []errors.Processor{printErr}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			e := errors.NewErrors(test.InitialProcessor)
			remove := e.Use(test.Processors...)
			remove()

			check(t, e.Debug())
		})
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		Name       string
		Processors []errors.Processor
		Error      error
	}{
		{"basic", []errors.Processor{new(multiErr).add}, errors.New("testing error")},
		{"two processors", []errors.Processor{new(multiErr).add, new(multiErr).log}, errors.New("testing errors")},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var (
				multiErr = new(multiErr)
				e        = errors.NewErrors(multiErr.add)
			)
			err := e.Process(test.Error)
			if err != test.Error {
				t.Fatalf("failed %s\nexpected = %s\nactual = %s", test.Name, test.Error, err)
			}

			check(t, multiErr.debug())
		})
	}
}

func TestProcessWith(t *testing.T) {
	tests := []struct {
		Name       string
		Processors []errors.Processor
		Error      error
	}{
		{"basic", []errors.Processor{new(multiErr).add}, errors.New("testing error")},
		{"two processors", []errors.Processor{new(multiErr).add, new(multiErr).log}, errors.New("testing errors")},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var (
				multiErr = new(multiErr)
				e        = errors.NewErrors()
			)
			err := e.ProcessWith(test.Error, test.Processors...)
			if err != test.Error {
				t.Fatalf("failed %s\nexpected = %s\nactual = %s", t.Name(), test.Error, err)
			}

			var str = strings.Join(
				[]string{multiErr.debug(), e.Debug()},
				"\n\n")
			check(t, str)
		})
	}
}

func check(t *testing.T, actual string) {
	if *golden {
		testutil.GoldenUpdate(t, []byte(actual))
	}

	var expected = string(testutil.GoldenGet(t))
	if actual != expected {
		t.Fatalf("failed %s\nexpected = %s\nactual = %s", t.Name(), expected, actual)
	}
}

func printErr(err error) {
	fmt.Printf("in printErr, err = %s", err)
}

func helloWorld(err error) {
	fmt.Printf("in helloWorld, err = %s", err)
}

type multiErr struct {
	errs   []error
	stdout bytes.Buffer
}

func (m *multiErr) add(err error) {
	m.errs = append(m.errs, err)
}

func (m *multiErr) log(err error) {
	fmt.Fprintf(&m.stdout, "in log, err = %s", err)
}

func (m *multiErr) debug() string {
	var debug string
	for _, err := range m.errs {
		debug += fmt.Sprint(err)
	}
	debug += fmt.Sprint(m.stdout.String())
	return debug
}
