package portapps

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/shlyo/kernal/pkg/utl"
)

// App represents an active app object
type App struct {
	ID   string
	Name string
	Args []string

	RootPath   string
	AppPath    string
	DataPath   string
	WorkingDir string
	Process    string

	logfile *os.File
	config  *Config
}

// NewWithCfg creates new app instance with an app config
func NewWithCfg(id string, name string, appcfg interface{}) (app *App, err error) {
	// Init
	app = &App{
		ID:   id,
		Name: name,
	}

	// Root path
	ex, err := os.Executable()
	app.RootPath, err = filepath.Abs(filepath.Dir(ex))

	// Load config
	err = app.loadConfig(appcfg)
	if appcfg != nil {
		err = mapstructure.Decode(app.config.App, appcfg)
	}

	// Set paths
	app.AppPath = utl.PathJoin(app.RootPath, "app")
	if app.config.Common.AppPath != "" {
		app.AppPath = app.config.Common.AppPath
	}
	app.DataPath = utl.PathJoin(app.RootPath, "data")
	app.WorkingDir = app.AppPath
	
	return app, nil
}

// Config returns app configuration
func (app *App) Config() *Config {
	return app.config
}

// Launch to execute the app with additional args
func (app *App) Launch(args []string) {
	jArgs := append(append(app.config.Common.Args, args...), app.Args...)
	execute := exec.Command(app.Process, jArgs...)
	execute.Dir = app.WorkingDir
	execute.Run()
}

// Close closes the app
func (app *App) Close() {
} 
