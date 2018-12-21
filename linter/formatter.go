package linter

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed).SprintFunc()
var prefix = red("error")

func ConsoleFormatter(out io.Writer, report Report) {
	violations := report.GetViolations()
	if len(violations) == 0 {
		fmt.Println("No errors found.")
	} else {
		for _, violation := range violations {
			fmt.Fprintf(out, "%v \"%v\" %v\n", prefix, violation.RuleName, violation.Failure)
		}
	}
}
