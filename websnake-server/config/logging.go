package config

import "github.com/fsufitch/websnake/websnake-server/log"

// ConfigureLogging configures custom logs according to the environment
func ConfigureLogging() error {
	config, err := GetConfig()
	if err != nil {
		return err
	}

	log.Debug.SetEnabled(config.DebugLogsEnabled)
	log.Debug.Print("Debug logging enabled.")
	return nil
}
