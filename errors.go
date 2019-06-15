//go:generate godocdown -o README.md

// Package errors is meant to manage the creation, logging, and various other
// operations performed after a process fails to operate as intended.
package errors

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"

	wrapped "github.com/pkg/errors"
)

// Errors is an object used to interact with the errors package's
// functionality. The various process dependencies are the fields of the Errors
// object.
type Errors struct {
	processors processors
}

// Processor is a function that accepts an error type and processes that
// object as it wishes.
type Processor func(error)

// NewErrors creates an instance of Errors object. The porcessors variadic
// argument specifies the Processor functions that are to be utilized in the
// processing of an error. The order of the processors variadic argument has no
// affect on the order the processors are executed.
func NewErrors(processors ...Processor) *Errors {
	e := new(Errors)
	e.With(processors...)
	return e
}

// Debug returns a string definition of the current set of Processors.
func (e *Errors) Debug() string {
	var debug string
	for _, p := range e.processors.get() {
		debug += fmt.Sprintf("%s\n", runtime.FuncForPC(reflect.ValueOf(p).Pointer()).Name())
	}
	return debug
}

// With overwrites the current set of Processors with a new set.
func (e *Errors) With(pSet ...Processor) {
	e.processors.with(pSet...)
}

// Use adds a set of Processors to the receiver's Processors set. A remove
// function is returned that may be called to remove the added Processors from
// the receiver's Processors set.
func (e *Errors) Use(processors ...Processor) func() {
	var pPtrs = make([]*Processor, len(processors))
	for i, p := range processors {
		pPtrs[i] = &p
	}
	e.processors.add(pPtrs...)
	return func() {
		e.processors.remove(pPtrs...)
	}
}

// Process returns nil if the err argument is nil. Process passes the specified
// err through a set of Processor functions owned by the reciever Errors
// object.
func (e *Errors) Process(err error) error {
	if err == nil {
		return nil
	}

	for _, p := range e.processors.get() {
		p(err)
	}
	return err
}

// ProcessWith returns nil if the err argument is nil. ProcessWith passes the
// specified err through the specified set(processors) of Processor functions.
func (e *Errors) ProcessWith(err error, processors ...Processor) error {
	if err == nil {
		return nil
	}

	for _, p := range processors {
		p(err)
	}
	return err
}

// Errorf serves as a wrapper for the github.com/pkg/errors.Errorf function.
func (e *Errors) Errorf(format string, args ...interface{}) error {
	return wrapped.Errorf(format, args...)
}

// Cause serves as a wrapper for the github.com/pkg/errors.Cause function.
func (e *Errors) Cause(err error) error {
	return wrapped.Cause(err)
}

// New serves as a wrapper for the github.com/pkg/errors.New function.
func (e *Errors) New(message string) error {
	return wrapped.New(message)
}

// WithMessage serves as a wrapper for the github.com/pkg/errors.WithMessage
// function.
func (e *Errors) WithMessage(err error, message string) error {
	return wrapped.WithMessage(err, message)
}

// WithStack serves as a wrapper for the github.com/pkg/errors.WithStack
// function.
func (e *Errors) WithStack(err error) error {
	return wrapped.WithStack(err)
}

// Wrap serves as a wrapper for the github.com/pkg/errors.Wrap function.
func (e *Errors) Wrap(err error, message string) error {
	return wrapped.Wrap(err, message)
}

// Wrapf serves as a wrapper for the github.com/pkg/errors.Wrapf function.
func (e *Errors) Wrapf(err error, format string, args ...interface{}) error {
	return wrapped.Wrapf(err, format, args...)
}

// processors is a set of Processor functions.
type processors struct {
	sync.RWMutex
	set []*Processor
}

// with replaces the current set of Processor objects on the Processor receiver
// with a new set of Processor objects. This is done in a concurrent-safe way.
func (ps *processors) with(pSet ...Processor) {
	var set = make([]*Processor, len(pSet))
	for i := range pSet {
		set[i] = &pSet[i]
	}
	ps.Lock()
	ps.set = set
	ps.Unlock()
}

// add appends a set of Processor objects to the Processor receiver in a
// concurrent-safe way.
func (ps *processors) add(pSet ...*Processor) {
	for _, p := range pSet {
		ps.Lock()
		ps.set = append(ps.set, p)
		ps.Unlock()
	}
}

// remove removes a set of Processor objects from the Processor receiver in a
// concurrent-safe way.
func (ps *processors) remove(pSet ...*Processor) {
	var hash = make(map[*Processor]int)
	ps.Lock()
	for i, p := range ps.set {
		hash[p] = i
	}
	for _, p := range pSet {
		index, ok := hash[p]
		if !ok {
			continue
		}

		var lastIndex = len(ps.set) - 1
		if index != lastIndex {
			ps.set[index] = ps.set[lastIndex]
			hash[ps.set[index]] = index
		}
		ps.set = ps.set[:lastIndex]
	}
	ps.Unlock()
}

// get retrieves a set of Processor objects from the Processors reciever in a
// concurrent-safe way.
func (ps *processors) get() []Processor {
	ps.RLock()
	defer ps.RUnlock()

	var processors = make([]Processor, len(ps.set))
	for i, p := range ps.set {
		processors[i] = *p
	}
	return processors
}

// global is an global instance of the Errors object.
var global = new(Errors)

// Debug serves as a global instance of errors.Errors.Debug().
func Debug() string {
	return global.Debug()
}

// With serves as a global instance of errors.Errors.With().
func With(processors ...Processor) {
	global.With(processors...)
}

// Use serves as a global instance of errors.Errors.Use().
func Use(processors ...Processor) func() {
	return global.Use(processors...)
}

// Process serves as a global instance of errors.Errors.Process().
func Process(err error) error {
	return global.Process(err)
}

// ProcessWith serves as a global instance of errors.Errors.ProcessWith().
func ProcessWith(err error, processors ...Processor) error {
	return global.ProcessWith(err, processors...)
}

// Errorf serves as a global instance of errors.Errors.Errorf().
func Errorf(format string, args ...interface{}) error {
	return global.Errorf(format, args...)
}

// Cause serves as a global instance of errors.Errors.Cause().
func Cause(err error) error {
	return global.Cause(err)
}

// New serves as a global instance of errors.Errors.New().
func New(message string) error {
	return global.New(message)
}

// WithMessage serves as a global instance of errors.Errors.WithMessage()
func WithMessage(err error, message string) error {
	return global.WithMessage(err, message)
}

// WithStack serves as a global instance of errors.Errors.WithStack().
func WithStack(err error) error {
	return global.WithStack(err)
}

// Wrap serves as a global instance of errors.Errors.Wrap().
func Wrap(err error, message string) error {
	return global.Wrap(err, message)
}

// Wrapf serves as a global instance of errors.Errors.Wrapf().
func Wrapf(err error, format string, args ...interface{}) error {
	return global.Wrapf(err, format, args...)
}
