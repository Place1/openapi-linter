package main

import (
	"github.com/place1/openapi-linter/linter"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	spec = kingpin.Arg("openapi-spec", "the path to an openapi spec yaml file").Required().String()
)

func main() {
	kingpin.Parse()
	linter.RunSpecLint(linter.Options{
		Spec: *spec,
	})
}
