package linter

import (
	"io/ioutil"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "reading file")
	}
	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, errors.Wrap(err, "parsing config")
	}

	for _, preset := range config.Presets {
		switch preset {
		case "standard":
			err := mergo.Merge(&config.Rules, &PresetStandard, mergo.WithOverride)
			if err != nil {
				return nil, errors.Wrap(err, "merging preset")
			}
		default:
			logrus.Warnf("no preset handler for %v implemented. skipping.", preset)
		}
	}

	return config, nil
}

type Config struct {
	Rules   Rules    `yaml:"rules,omitempty"`
	Presets []string `yaml:"presets,omitempty"`
}

type Rules struct {
	NoEmptyDescriptions          *NoEmptyDescriptionsOpts          `yaml:"noEmptyDescriptions,omitempty"`
	NoEmptyOperationID           bool                              `yaml:"noEmptyOperationID,omitempty"`
	SlashTerminatedPaths         bool                              `yaml:"slashTerminatedPaths,omitempty"`
	RequireOperationTags         bool                              `yaml:"requireOperationTags,omitempty"`
	PathNamingConvention         *PathNamingConventionOpts         `yaml:"pathNamingConvention,omitempty"`
	DefinitionNamingConvention   *DefinitionNamingConventionOpts   `yaml:"definitionNamingConvention,omitempty"`
	PropertyNamingConvention     *PropertyNamingConventionOpts     `yaml:"propertyNamingConvention,omitempty"`
	ParameterNamingConvention    *ParameterNamingConventionOpts    `yaml:"parameterNamingConvention,omitempty"`
	OperationTagNamingConvention *OperationTagNamingConventionOpts `yaml:"operationTagNamingConvention,omitempty"`
}

type NoEmptyDescriptionsOpts struct {
	IgnoreProperties bool `yaml:"ignoreProperties,omitempty"`
	IgnoreOperations bool `yaml:"ignoreOperations,omitempty"`
	IgnoreParameters bool `yaml:"ignoreParameters,omitempty"`
}

type PathNamingConventionOpts struct {
	Convention NamingConvention `yaml:"convention"`
}

type DefinitionNamingConventionOpts struct {
	Convention NamingConvention `yaml:"convention"`
}

type PropertyNamingConventionOpts struct {
	Convention NamingConvention `yaml:"convention"`
}

type ParameterNamingConventionOpts struct {
	Convention NamingConvention `yaml:"convention"`
}

type OperationTagNamingConventionOpts struct {
	Convention NamingConvention `yaml:"convention"`
}

var PresetStandard = Rules{
	NoEmptyDescriptions: &NoEmptyDescriptionsOpts{
		IgnoreProperties: true,
		IgnoreParameters: true,
	},
	NoEmptyOperationID:   true,
	RequireOperationTags: true,
	DefinitionNamingConvention: &DefinitionNamingConventionOpts{
		Convention: PascalCase,
	},
	OperationTagNamingConvention: &OperationTagNamingConventionOpts{
		Convention: CamelCase,
	},
}
