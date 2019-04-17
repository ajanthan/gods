package model

type Resource struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
	Query  Query  `yaml:"query"`
}
