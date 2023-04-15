package serviceYaml

type ServiceYAML struct {
	Version     string     `yaml:"version"`
	Language    string     `yaml:"language"`
	ServiceName string     `yaml:"name"`
	Contract    Contract   `yaml:"contract"`
	Dependency  Dependency `yaml:"dependencies"`
}

type Contract struct {
	OutputGrst string   `yaml:"output-grst"`
	ProtoFiles []string `yaml:"proto-files"`
}

type Dependency struct {
	OutputGrst string   `yaml:"output-grst"`
	Services   []string `yaml:"services"`
}
