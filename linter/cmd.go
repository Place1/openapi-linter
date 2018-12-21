package linter

import (
	"log"
	"os"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
	"github.com/place1/openapi-linter/core"
)

type Options struct {
	Spec   string
	Config string
}

func RunSpecLint(options Options) {
	document, err := loads.Spec(options.Spec)
	if err != nil {
		log.Fatalln(err)
	}

	config, err := LoadConfig(options.Config)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := &RuleContext{
		Config:   *config,
		Report:   *NewReport(),
		Analyzer: *analysis.New(document.Spec()),
		Walk: func(visitor core.DocumentVisitor) {
			core.Walk(document, core.NodeData{}, visitor)
		},
	}

	Naming(ctx)
	RequireOperationTags(ctx)
	NoEmptyOperationID(ctx)
	NoEmptyDescriptions(ctx)
	SlashTerminatedPaths(ctx)
	NoUnusedDefinitions(ctx)
	NoDuplicateOperationIDs(ctx)

	// validator := validate.NewSpecValidator(document.Schema(), strfmt.Default)
	// validator.SetContinueOnErrors(true)
	// result, _ := validator.Validate(document)
	// for _, specError := range result.Errors {
	// 	logrus.Error(specError.Error())
	// }
	// for _, warning := range result.Warnings {
	// 	logrus.Warn(warning.Error())
	// }

	ConsoleFormatter(os.Stdout, ctx.Report)
}
