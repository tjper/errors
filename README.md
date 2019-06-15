# errors
--
    import "."

Package errors is meant to manage the creation, logging, and various other
operations performed after a process fails to operate as intended.

## Usage

#### func  Cause

```go
func Cause(err error) error
```
Cause serves as a global instance of errors.Errors.Cause().

#### func  Debug

```go
func Debug() string
```
Debug serves as a global instance of errors.Errors.Debug().

#### func  Errorf

```go
func Errorf(format string, args ...interface{}) error
```
Errorf serves as a global instance of errors.Errors.Errorf().

#### func  New

```go
func New(message string) error
```
New serves as a global instance of errors.Errors.New().

#### func  Process

```go
func Process(err error) error
```
Process serves as a global instance of errors.Errors.Process().

#### func  ProcessWith

```go
func ProcessWith(err error, processors ...Processor) error
```
ProcessWith serves as a global instance of errors.Errors.ProcessWith().

#### func  Use

```go
func Use(processors ...Processor) func()
```
Use serves as a global instance of errors.Errors.Use().

#### func  With

```go
func With(processors ...Processor)
```
With serves as a global instance of errors.Errors.With().

#### func  WithMessage

```go
func WithMessage(err error, message string) error
```
WithMessage serves as a global instance of errors.Errors.WithMessage()

#### func  WithStack

```go
func WithStack(err error) error
```
WithStack serves as a global instance of errors.Errors.WithStack().

#### func  Wrap

```go
func Wrap(err error, message string) error
```
Wrap serves as a global instance of errors.Errors.Wrap().

#### func  Wrapf

```go
func Wrapf(err error, format string, args ...interface{}) error
```
Wrapf serves as a global instance of errors.Errors.Wrapf().

#### type Errors

```go
type Errors struct {
}
```

Errors is an object used to interact with the errors package's functionality.
The various process dependencies are the fields of the Errors object.

#### func  NewErrors

```go
func NewErrors(processors ...Processor) *Errors
```
NewErrors creates an instance of Errors object. The porcessors variadic argument
specifies the Processor functions that are to be utilized in the processing of
an error. The order of the processors variadic argument has no affect on the
order the processors are executed.

#### func (*Errors) Cause

```go
func (e *Errors) Cause(err error) error
```
Cause serves as a wrapper for the github.com/pkg/errors.Cause function.

#### func (*Errors) Debug

```go
func (e *Errors) Debug() string
```
Debug returns a string definition of the current set of Processors.

#### func (*Errors) Errorf

```go
func (e *Errors) Errorf(format string, args ...interface{}) error
```
Errorf serves as a wrapper for the github.com/pkg/errors.Errorf function.

#### func (*Errors) New

```go
func (e *Errors) New(message string) error
```
New serves as a wrapper for the github.com/pkg/errors.New function.

#### func (*Errors) Process

```go
func (e *Errors) Process(err error) error
```
Process returns nil if the err argument is nil. Process passes the specified err
through a set of Processor functions owned by the reciever Errors object.

#### func (*Errors) ProcessWith

```go
func (e *Errors) ProcessWith(err error, processors ...Processor) error
```
ProcessWith returns nil if the err argument is nil. ProcessWith passes the
specified err through the specified set(processors) of Processor functions.

#### func (*Errors) Use

```go
func (e *Errors) Use(processors ...Processor) func()
```
Use adds a set of Processors to the receiver's Processors set. A remove function
is returned that may be called to remove the added Processors from the
receiver's Processors set.

#### func (*Errors) With

```go
func (e *Errors) With(pSet ...Processor)
```
With overwrites the current set of Processors with a new set.

#### func (*Errors) WithMessage

```go
func (e *Errors) WithMessage(err error, message string) error
```
WithMessage serves as a wrapper for the github.com/pkg/errors.WithMessage
function.

#### func (*Errors) WithStack

```go
func (e *Errors) WithStack(err error) error
```
WithStack serves as a wrapper for the github.com/pkg/errors.WithStack function.

#### func (*Errors) Wrap

```go
func (e *Errors) Wrap(err error, message string) error
```
Wrap serves as a wrapper for the github.com/pkg/errors.Wrap function.

#### func (*Errors) Wrapf

```go
func (e *Errors) Wrapf(err error, format string, args ...interface{}) error
```
Wrapf serves as a wrapper for the github.com/pkg/errors.Wrapf function.

#### type Processor

```go
type Processor func(error)
```

Processor is a function that accepts an error type and processes that object as
it wishes.
