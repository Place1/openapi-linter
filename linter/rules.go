package linter

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/place1/openapi-linter/core"
	"github.com/place1/openapi-linter/utils"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/jsonpointer"
	"github.com/go-openapi/spec"
)

type RuleContext struct {
	Config   Config
	Analyzer analysis.Spec
	Report   Report
	Walk     func(visitor core.DocumentVisitor)
}

func NoEmptyDescriptions(ctx *RuleContext) {
	if ctx.Config.Rules.NoEmptyDescriptions == nil {
		return
	}

	ctx.Walk(func(node interface{}, data core.NodeData) {
		switch node := node.(type) {
		case *spec.Operation:
			if node.Description == "" {
				ctx.Report.AddViolation(RuleViolation{
					RuleName: "",
					Failure:  fmt.Sprintf(`operation "%v" must have a description`, data.Ref),
				})
			}
		}
	})
}

func NoEmptyOperationID(ctx *RuleContext) {
	if ctx.Config.Rules.NoEmptyOperationIDs == false {
		return
	}

	ctx.Walk(func(node interface{}, data core.NodeData) {
		switch node := node.(type) {
		case *spec.Operation:
			if node.ID == "" {
				ctx.Report.AddViolation(RuleViolation{
					RuleName: "",
					Failure:  fmt.Sprintf(`operation "%v" must have an operationId`, data.Ref),
				})
			}
		}
	})
}

func SlashTerminatedPaths(ctx *RuleContext) {
	if ctx.Config.Rules.SlashTerminatedPaths == nil {
		return
	}

	ctx.Walk(func(node interface{}, data core.NodeData) {
		switch node.(type) {
		case *spec.PathItem:
			if *ctx.Config.Rules.SlashTerminatedPaths && strings.HasSuffix(data.Key, "/") {
				ctx.Report.AddViolation(RuleViolation{
					RuleName: "",
					Failure:  fmt.Sprintf(`path "%v" must end with a slash`, data.Ref),
				})
			}
		}
	})
}

type NamingConvention = string

const (
	PascalCase NamingConvention = "PascalCase"
	SnakeCase  NamingConvention = "SnakeCase"
	CamelCase  NamingConvention = "CamelCase"
	KebabCase  NamingConvention = "KebabCase"
)

func GetNamingChecker(convention NamingConvention) func(string) bool {
	var checker func(string) bool
	switch convention {
	case PascalCase:
		checker = utils.IsPascalCase
	case CamelCase:
		checker = utils.IsCamelCase
	case SnakeCase:
		checker = utils.IsSnakeCase
	case KebabCase:
		checker = utils.IsKebabCase
	}
	return checker
}

func Naming(ctx *RuleContext) {
	if ctx.Config.Rules.Naming == nil {
		return
	}

	ctx.Walk(func(node interface{}, data core.NodeData) {
		switch node := node.(type) {
		case *spec.PathItem:
			if ctx.Config.Rules.Naming.Paths != "" {
				checker := GetNamingChecker(ctx.Config.Rules.Naming.Paths)
				pathSegments := strings.Split(data.Key, "/")
				for _, segment := range pathSegments {
					if segment != "" && !strings.HasPrefix(segment, "{") && !checker(segment) {
						ctx.Report.AddViolation(RuleViolation{
							RuleName: "Naming",
							Failure:  fmt.Sprintf(`path "%v" must follow the %v naming convention`, data.Key, ctx.Config.Rules.Naming.Paths),
						})
					}
				}
			}

		case *spec.Definitions:
			if ctx.Config.Rules.Naming.Definitions != "" {
				checker := GetNamingChecker(ctx.Config.Rules.Naming.Definitions)
				for name := range *node {
					if !checker(name) {
						ctx.Report.AddViolation(RuleViolation{
							RuleName: "",
							Failure:  fmt.Sprintf(`definition "%v" must follow the %v naming convention`, name, ctx.Config.Rules.Naming.Definitions),
						})
					}
				}
			}

		case *spec.Schema:
			if ctx.Config.Rules.Naming.Properties != "" {
				checker := GetNamingChecker(ctx.Config.Rules.Naming.Properties)
				for name := range node.Properties {
					if !checker(name) {
						ctx.Report.AddViolation(RuleViolation{
							RuleName: "Naming.Properties",
							Failure:  fmt.Sprintf(`property "%v" must follow the %v naming convention`, name, ctx.Config.Rules.Naming.Properties),
						})
					}
				}
			}

		case *spec.Parameter:
			if ctx.Config.Rules.Naming.Parameters != "" {
				checker := GetNamingChecker(ctx.Config.Rules.Naming.Parameters)
				if !checker(node.Name) {
					ctx.Report.AddViolation(RuleViolation{
						RuleName: "",
						Failure:  fmt.Sprintf(`parameter "%v" must follow the %v naming convention`, data.Ref, ctx.Config.Rules.Naming.Parameters),
					})
				}
			}

		case *spec.Operation:
			if ctx.Config.Rules.Naming.Operations != "" {
				checker := GetNamingChecker(ctx.Config.Rules.Naming.Operations)
				if !checker(node.ID) {
					ctx.Report.AddViolation(RuleViolation{
						RuleName: "",
						Failure:  fmt.Sprintf(`operation with id "%v" must follow the %v naming convention`, node.ID, ctx.Config.Rules.Naming.Operations),
					})
				}
			}

			if ctx.Config.Rules.Naming.Tags != "" {
				checker := GetNamingChecker(ctx.Config.Rules.Naming.Tags)
				for i, tag := range node.Tags {
					if !checker(tag) {
						ctx.Report.AddViolation(RuleViolation{
							RuleName: "",
							Failure:  fmt.Sprintf(`operation tag "%v" must follow the %v naming convention`, path.Join(data.Ref, strconv.Itoa(i)), ctx.Config.Rules.Naming.Tags),
						})
					}
				}
			}
		}
	})
}

func RequireOperationTags(ctx *RuleContext) {
	if ctx.Config.Rules.NoEmptyTags == false {
		return
	}
	ctx.Walk(func(node interface{}, data core.NodeData) {
		switch node := node.(type) {
		case *spec.Operation:
			if len(node.Tags) == 0 {
				ctx.Report.AddViolation(RuleViolation{
					RuleName: "",
					Failure:  fmt.Sprintf(`operation "%v" must have at least 1 tag`, data.Ref),
				})
			}
		}
	})
}

func NoUnusedDefinitions(ctx *RuleContext) {
	if ctx.Config.Rules.NoUnusedDefinitions == false {
		return
	}

	expected := map[string]bool{}
	references := ctx.Analyzer.AllDefinitionReferences()

	ctx.Walk(func(node interface{}, data core.NodeData) {
		switch node := node.(type) {
		case *spec.Definitions:
			for name := range *node {
				expected["#/definitions/"+jsonpointer.Escape(name)] = true
			}

			for _, name := range references {
				if _, ok := expected[name]; ok {
					delete(expected, name)
				}
			}

			for name := range expected {
				ctx.Report.AddViolation(RuleViolation{
					RuleName: "NoUnusedDefinitions",
					Failure:  fmt.Sprintf("definition %v is unused", name),
				})
			}
		}
	})
}

func NoDuplicateOperationIDs(ctx *RuleContext) {
	if ctx.Config.Rules.NoDuplicateOperationIDs == false {
		return
	}

	existing := map[string]bool{}

	ctx.Walk(func(node interface{}, data core.NodeData) {
		switch node := node.(type) {
		case *spec.Operation:
			if node.ID != "" {
				if _, ok := existing[node.ID]; !ok {
					// first time seeing this operation id
					existing[node.ID] = true
				} else {
					// we've already seen it, it's a duplicate
					ctx.Report.AddViolation(RuleViolation{
						RuleName: "NoDuplicateOperationIDs",
						Failure:  fmt.Sprintf("operation id \"%v\" is a duplicate", node.ID),
					})
				}
			}
		}
	})
}

func NoMissingRequiredProperties(ctx *RuleContext) {
	if ctx.Config.Rules.NoMissingRequiredProperties == false {
		return
	}

	ctx.Walk(func(node interface{}, data core.NodeData) {
		switch node := node.(type) {
		case *spec.Schema:
			for _, requiredProperty := range node.Required {
				if _, ok := node.Properties[requiredProperty]; !ok {
					ctx.Report.AddViolation(RuleViolation{
						RuleName: "NoMissingRequiredProperties",
						Failure:  fmt.Sprintf("property \"%v\" is listed as required but is not defined under \"properties\"", requiredProperty),
					})
				}
			}
		}
	})
}
