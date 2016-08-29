package config

// ParameterContainer holds a list of parameters
type ParameterContainer struct {
	List []*Parameter
}

// Get returns the first parameter with the given name
func (pc *ParameterContainer) Get(name string) *Parameter {
	for _, p := range pc.List {
		if p.Key == name {
			return p
		}
	}

	return &Parameter{}
}

// Value returns the value of the first parameter with the given name
func (pc *ParameterContainer) Value(name string) string {
	p := pc.Get(name)
	if p.Value == "" {
		return p.Default
	}
	return p.Value
}
