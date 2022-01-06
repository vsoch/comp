package options

import (
	"fmt"
)

// A list of options for a backend or plugin
type Options []Option

// An option can come from a config, the environment, or a command line variable
// TODO have a means for any backend/plugin to define opts in the environment
type Option struct {
	Name     string
	Help     string
	Required bool
}

// Set the default values for the options
// TODO here maybe we should include parsing enviroment from config?
// @alecbcs we should figure out how to best integrate config with backends/options generally
func (options Options) Init() {
	for _, option := range options {
		fmt.Println("TODO set value for option %s", option)
	}
}

// Get a named option
func (options Options) Get(name string) *Option {
	for _, option := range options {
		if option.Name == name {
			return &option
		}
	}
	return nil
}
