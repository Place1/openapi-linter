package linter

import (
	"io/ioutil"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find config file %v", path)
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, errors.Wrap(err, "parsing config")
	}

	return config, nil
}

type Config struct {
	Rules Rules `yaml:"rules,omitempty"`
}

type Rules struct {
	Naming                  *NamingOpts              `yaml:"naming,omitempty"`
	NoEmptyDescriptions     *NoEmptyDescriptionsOpts `yaml:"noEmptyDescriptions,omitempty"`
	NoEmptyOperationIDs     bool                     `yaml:"noEmptyOperationIDs,omitempty"`
	SlashTerminatedPaths    *bool                    `yaml:"slashTerminatedPaths,omitempty"`
	NoEmptyTags             bool                     `yaml:"noEmptyTags,omitempty"`
	NoUnusedDefinitions     bool                     `yaml:"noUnusedDefinitions,omitempty"`
	NoDuplicateOperationIDs bool                     `yaml:"noDuplicateOperationIDs,omitempty"`
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
