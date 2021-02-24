package app

import (
	"fmt"
	"io/ioutil"
)

// ExportStateToJSON util function to export the app state to JSON
func ExportStateToJSON(app *StraightedgeApp, path string) error {
	fmt.Println("exporting app state...")
	exportedApp, err := app.ExportAppStateAndValidators(false, nil)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, []byte(exportedApp.AppState), 0644)
}
