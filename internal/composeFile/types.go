package composeFile

type ComposeService struct {
	Build struct {
		Context    string `yaml:"context"`
		Dockerfile string `yaml:"dockerfile"`
		Target     string `yaml:"target"`
	} `yaml:"build"`
	Image string `yaml:"image"`
}

type ComposeFile struct {
	Services map[string]ComposeService `yaml:"services"`
}