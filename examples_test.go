package errors_test

import (
	"fmt"

	"github.com/ShaleApps/errors"
)

func ExampleProcess() {
	var log = func(err error) {
		fmt.Println(err)
	}
	var alert = func(err error) {
		fmt.Println("HEY this error just occured")
	}

	var (
		e   = errors.NewErrors(log, alert)
		err = errors.New("failed to do something")
	)
	e.Process(err)

	// Output:
	// failed to do something
	// HEY this error just occured
}

func ExampleProcessWith() {
	var log = func(err error) {
		fmt.Println(err)
	}
	var alert = func(err error) {
		fmt.Println("HEY this error just occured")
	}

	var (
		e   = errors.NewErrors(log)
		err = errors.New("failed to do something")
	)
	e.ProcessWith(err, alert)

	// Output:
	// HEY this error just occured
}

func ExampleUse() {
	var log = func(err error) {
		fmt.Println(err)
	}
	var alert = func(err error) {
		fmt.Println("HEY this error just occured")
	}

	var (
		e   = errors.NewErrors(log)
		err = errors.New("failed to do something")
	)
	var remove = e.Use(alert)

	e.Process(err)
	remove()
	e.Process(err)

	// Output:
	// failed to do something
	// HEY this error just occured
	// failed to do something
}
