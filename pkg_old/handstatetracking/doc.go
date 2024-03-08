/*
File: doc.go, version history:

v0.1	2021/05/25	Tomasz Tudrej

License info

Copyright info

First line of documentation text in the doc.go file
Any additional documentation for the package goes in this file.

The "package <package name>"" line, MUST follow immediately after the end of comment,
with NO empty lines in between, in order to be used by "go doc"

So, it seems that "go doc" will search the entire package directory for files containing the
"package <pkg name>" declaration, and grab the commnets, if any, that prodeede those lines
immediattely, with not empty lines, and display them as it's output.

*/
package handstatetracking
