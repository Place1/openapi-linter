package main

import (
	"github.com/place1/openapi-linter/linter"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	spec   = kingpin.Arg("openapi-spec", "a path to an openapi yaml spec").Required().String()
	config = kingpin.Flag("config", "a path to a config file").Default("openapi-linter.yaml").String()
)

func main() {
	kingpin.Parse()
	linter.RunSpecLint(linter.Options{
		Spec: *spec,
	})
}
