package sdk

// Function represents a single-invocation function that responds to events,
// performs actions, or is exposed via the API.
type Function struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description,omitempty"`
	Runtime     string `yaml:"runtime"`
}
