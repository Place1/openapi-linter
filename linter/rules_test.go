package linter

import (
	"encoding/json"
	"testing"

	"github.com/place1/openapi-linter/utils"
	"github.com/stretchr/testify/require"

	"github.com/go-openapi/loads"
	"github.com/place1/openapi-linter/core"

	"github.com/imdario/mergo"
)

type TestContext struct {
	Spec   string
	Config Config
}

func CreateContext(custom TestContext) *RuleContext {
	document, err := loads.Analyzed(json.RawMessage([]byte(custom.Spec)), "")
	if err != nil {
		panic(err)
	}

	ctx := RuleContext{
		Walk: func(visitor core.DocumentVisitor) {
			core.Walk(document, core.NodeData{}, visitor)
		},
	}

	err = mergo.Merge(&ctx.Config, &custom.Config, mergo.WithOverride)
	if err != nil {
		panic(err)
	}

	return &ctx
}

func TestPascalCaseNaming(t *testing.T) {
	require := require.New(t)
	ctx := CreateContext(TestContext{
		Spec: utils.Yaml(`
			paths:
				/MyPath/{id}/ThatIsCamelCase/:
					hello: world
		`),
		Config: Config{
			Rules: Rules{
				Naming: &NamingOpts{
					Paths: "PascalCase",
				},
			},
		},
	})

	Naming(ctx)

	require.Empty(ctx.Report.GetViolations())
}
