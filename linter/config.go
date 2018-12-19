package linter

import (
	"io/ioutil"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

var PresetStandard = Rules{
	NoEmptyDescriptions: &NoEmptyDescriptionsOpts{
		Operations: true,
	},
	NoEmptyOperationIDs: true,
	NoEmptyTags:         true,
	Naming: &NamingOpts{
		Tags:        PascalCase,
		Operations:  CamelCase,
		Definitions: PascalCase,
	},
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find config file")
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
	Naming               *NamingOpts              `yaml:"naming,omitempty"`
	NoEmptyDescriptions  *NoEmptyDescriptionsOpts `yaml:"noEmptyDescriptions,omitempty"`
	NoEmptyOperationIDs  bool                     `yaml:"noEmptyOperationIDs,omitempty"`
	SlashTerminatedPaths bool                     `yaml:"slashTerminatedPaths,omitempty"`
	NoEmptyTags          bool                     `yaml:"noEmptyTags,omitempty"`
}

type NoEmptyDescriptionsOpts struct {
	Operations bool `yaml:"operations,omitempty"`
	Properties bool `yaml:"properties,omitempty"`
	Parameters bool `yaml:"parameters,omitempty"`
}

type NamingOpts struct {
	Paths       NamingConvention `yaml:"paths,omitempty"`
	Tags        NamingConvention `yaml:"tags,omitempty"`
	Operations  NamingConvention `yaml:"operation,omitempty"`
	Parameters  NamingConvention `yaml:"parameters,omitempty"`
	Definitions NamingConvention `yaml:"definitions,omitempty"`
	Properties  NamingConvention `yaml:"properties,omitempty"`
}
