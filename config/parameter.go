package config

// Parameter is a key/value object
type Parameter struct {
	Key         string
	Value       string
	Default     string
	Description string
}

// Set the parameter value
func (p *Parameter) Set(value string) {
	p.Value = value
}
