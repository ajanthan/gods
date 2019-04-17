package model

type Query struct {
	//Ref    string   `yaml:"$ref"`
	//ID     string   `yaml:"id"`
	SQL    string   `yaml:"sql"`
	Params []Params `yaml:"params,omitempty"`
	Result Result   `yaml:"result,omitempty"`
}

type Params struct {
	Name    string `yaml:"name"`
	SQLType string `yaml:"sqlType"`
	Ordinal int    `yaml:"ordinal"`
}

type Result struct {
	Type   string  `yaml:"type,omitempty"`
	Fields []Field `yaml:"schema,omitempty"`
}

type Field struct {
	Name   string `yaml:"name,omitempty"`
	Column string `yaml:"column,omitempty"`
}
