// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package instrument

// ExecFn is an executable function that can be instrumented with a Call.
type ExecFn func() error

// FilterFn takes in an error and, if the error does not actually indicate
// a failure, returns true. If the error does actually indicate failure, the
// function returns false.
type FilterFn func(err error) bool

// Call allows tracking the successes, errors, and timing of functions.
type Call interface {
	// Exec executes a function and records whether it succeeded or
	// failed, and the amount of time that it took.
	Exec(f ExecFn) error

	// Exec executes a function and records whether it succeeded or
	// failed, and the amount of time that it took. ExecWithFilter
	// will always return the error associted with the called function.
	// However, it will decide whether or not classify that error as
	// a success or failure based on the provided SuccessFilter.
	ExecWithFilter(f ExecFn, filter FilterFn) error
}