// Package debugging comment string
package debugging

import "runtime"

func ThisFunc() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	// file, line := f.FileLine(pc[0])
	// fmt.Printf("%s\n", f.Name())
	return f.Name()
}
