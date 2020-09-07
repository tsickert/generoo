package generoo

type Configuration struct {
	Metadata
	Constants []Constant
	Inputs    []Input
}

type Metadata struct {
	Version     string
	Name        string
	Description string
}

type Constant struct {
	Name  string
	Value string
}

type Input struct {
	Name        string
	Text        string
	Type        string
	Default     string
	Validations []Validation
	FollowUps   []Input
}

type Validation struct {
	Type  string
	Value string
}
