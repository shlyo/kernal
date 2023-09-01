package portapps

import (
	"fmt"
	"os"

	"github.com/shlyo/kernal/pkg/utl"
	"gopkg.in/yaml.v3"
)

// Config holds portapp configuration details
type Config struct {
	Common Common      `yaml:"common" mapstructure:"common"`
	App    interface{} `yaml:"app,omitempty" mapstructure:"app"`
}

// Common holds common configuration data
type Common struct {
	DisableLog bool              `yaml:"disable_log" mapstructure:"disable_log"`
	Args       []string          `yaml:"args" mapstructure:"args"`
	AppPath    string            `yaml:"app_path" mapstructure:"app_path"`
}

// loadConfig load common and app configuration
func (app *App) loadConfig(appcfg interface{}) (err error) {
	cfgPath := utl.PathJoin(app.RootPath, fmt.Sprintf("%s.yml", app.ID))
	app.config = &Config{
		Common: Common{
			DisableLog: true,
			Args:       []string{},
			AppPath:    "",
		},
		App: appcfg,
	}

	// Write sample config
	raw, err := yaml.Marshal(app.config)

	// Read config
	raw, err = os.ReadFile(cfgPath)

	return yaml.Unmarshal(raw, &app.config)
}
