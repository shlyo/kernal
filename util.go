package portapps

import (
	"github.com/shlyo/kernal/pkg/utl"
)

// ElectronAppPath returns the app electron path
func (app *App) ElectronAppPath() string  {
	electronAppFolder, err := utl.FindElectronAppFolder("app-", app.AppPath)
	if err != nil {
	}
	return utl.PathJoin(app.AppPath, electronAppFolder)
}
