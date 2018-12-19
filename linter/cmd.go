package linter

import (
	"log"
	"sync"

	"github.com/place1/openapi-linter/core"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	"github.com/sirupsen/logrus"
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

	_, err = LoadConfig(options.Config)
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		core.Walk(document, core.NodeData{}, OperationTagNamingConvention(KebabCase))
		core.Walk(document, core.NodeData{}, RequireOperationTags())
		core.Walk(document, core.NodeData{}, PathNamingConvention(KebabCase))
		core.Walk(document, core.NodeData{}, ParameterNamingConvention(KebabCase))
		core.Walk(document, core.NodeData{}, DefinitionNamingConvention(KebabCase))
		core.Walk(document, core.NodeData{}, PropertyNamingConvention(KebabCase))
		core.Walk(document, core.NodeData{}, NoEmptyOperationID())
		core.Walk(document, core.NodeData{}, NoEmptyDescriptions())
		core.Walk(document, core.NodeData{}, SlashTerminatedPaths())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		validator := validate.NewSpecValidator(document.Schema(), strfmt.Default)
		validator.SetContinueOnErrors(true)
		result, _ := validator.Validate(document)
		for _, specError := range result.Errors {
			logrus.Error(specError.Error())
		}
		for _, warning := range result.Warnings {
			logrus.Warn(warning.Error())
		}
	}()

	wg.Wait()
}
