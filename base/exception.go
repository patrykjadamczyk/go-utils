package base

import (
	"fmt"
	"runtime"
)

type ExceptionCategory string

type Exception struct {
	Message  string
	Category ExceptionCategory
}

func (e Exception) GuardInit() {
	AssertThatTyped2[error, Exception]().ImplementsT1(e)
}

func (e *Exception) Init(category ExceptionCategory, msg string) {
	e.Category = category
	e.Message = msg
}

func (e Exception) Error() string {
	return fmt.Sprintf("%s: %s", e.Category, e.Message)
}

type ExtendedException struct {
	Exception
	Stacktrace string
	Notes []string
}

func (e ExtendedException) GuardInit() {
	AssertThatTyped2[error, ExtendedException]().ImplementsT1(e)
}

func (e *ExtendedException) Init(category ExceptionCategory, msg string) {
	e.Category = category
	e.Message = msg
	e.Stacktrace = e.fillStackTrace()
	e.Notes = make([]string, 0)
}

func (e *ExtendedException) AddNote(note string) {
	e.Notes = append(e.Notes, note)
}

func (e ExtendedException) fillStackTrace() string {
	maxFrames := 10
	skipFrames := 3
	resultStacktrace := ""
	pcs := make([]uintptr, maxFrames)
	numFrames := runtime.Callers(skipFrames, pcs)
	frames := runtime.CallersFrames(pcs[:numFrames])

	for {
		frame, more := frames.Next()
		resultStacktrace = fmt.Sprintf(
			"%s\n%s()\n  At %s:%d",
			resultStacktrace,
			frame.Function,
			frame.File,
			frame.Line,
		)

		if !more {
			break
		}
	}

	return resultStacktrace
}

type SubCategoryExtendedException struct {
	ExtendedException
}

func (e SubCategoryExtendedException) GuardInit() {
	AssertThatTyped2[error, SubCategoryExtendedException]().ImplementsT1(e)
}

func (e *SubCategoryExtendedException) Init(category ExceptionCategory, subcategory ExceptionCategory, msg string) {
	e.Category = ExceptionCategory(fmt.Sprintf("%s: %s", category, subcategory))
	e.Message = msg
	e.Stacktrace = e.fillStackTrace()
	e.Notes = make([]string, 0)
}
