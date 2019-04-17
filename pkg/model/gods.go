package model

type Gods struct {
	Version    string     `yaml:"gods"`
	Name       string     `yaml:"name"`
	BaseURI    string     `yaml:"baseURI"`
	Datasource Datasource `yaml:"datasource"`
	Resources  []Resource `yaml:"resources"`
	//Queries    []Query    `yaml:"queries"`
}
