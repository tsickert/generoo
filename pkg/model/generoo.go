package model

type Metadata struct {
	Id          string
	Name        string
	Version     string
	Description string
}

type Generoo struct {
	Metadata
	Constants []Constant
	Inputs    []Input
}

type Constant struct {
}

type Input struct {
}
