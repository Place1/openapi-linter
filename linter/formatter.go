package linter

import (
	"fmt"
	"io"
	"sort"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed).SprintFunc()
var prefix = red("error")

func ConsoleFormatter(out io.Writer, report Report) {
	violations := report.GetViolations()
	sort.Slice(violations, func(i int, j int) bool {
		return violations[i].Ref < violations[j].Ref
	})
	if len(violations) == 0 {
		fmt.Println("No errors found.")
	} else {
		for _, violation := range violations {
			fmt.Fprintf(out, "%v \"%v\" \"%v\" %v\n", prefix, violation.Ref, violation.RuleName, violation.Failure)
		}
	}
}
