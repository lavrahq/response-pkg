package sdk

// App is the app that various objects belong to.
type App struct {
	Name        string
	Description string

	Functions []*Function
}

// CreateApp initializes and returns a new App instance.
func CreateApp(name string, description string) *App {
	return &App{
		Name:        name,
		Description: description,
	}
}

// SetFunctionPath sets the path in the repository that the installer
// can find the functions.
func (a *App) SetFunctionPath() {

}
