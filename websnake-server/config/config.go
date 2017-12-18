package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fsufitch/websnake/websnake-server/log"
)

// ServerConfiguration is a struct encompassing arguments the server is to
// be configured with
type ServerConfiguration struct {
	Port             int
	APIHost          string
	UIStaticPath     string
	CacheTTL         time.Duration
	DebugLogsEnabled bool
}

func (s ServerConfiguration) getServeAddress() string {
	return fmt.Sprintf(":%d", s.Port)
}

var cachedConfig *ServerConfiguration

func buildConfig() (config *ServerConfiguration, err error) {
	log.Info.Print("Building environment configuration...")
	config = &ServerConfiguration{}

	if portStr, ok := os.LookupEnv("PORT"); ok {
		config.Port, err = strconv.Atoi(portStr)
	} else {
		config.Port = 8080
		log.Error.Printf("PORT env variable not found; using default: %v", config.Port)
	}
	if err != nil {
		return
	}
	log.Info.Printf("Port = %v", config.Port)

	if apiHost, ok := os.LookupEnv("API_HOST"); ok {
		config.APIHost = apiHost
	} else {
		err = errors.New("API_HOST env variable not found")
		return
	}
	log.Info.Printf("APIHost = %v", config.APIHost)

	if uiStaticPath, ok := os.LookupEnv("UI_STATIC_PATH"); ok {
		config.UIStaticPath = uiStaticPath
	} else {
		err = errors.New("UI_STATIC_PATH env variable not found")
	}
	if err != nil {
		return
	}
	log.Info.Printf("UIStaticPath = %v", config.UIStaticPath)

	if cacheTTLStr, ok := os.LookupEnv("CACHE_TTL"); ok {
		var seconds int
		seconds, err = strconv.Atoi(cacheTTLStr)
		config.CacheTTL = time.Duration(seconds) * time.Second
	} else {
		config.CacheTTL = 1 * time.Second
		log.Error.Printf("PROXY_TTL env variable not found; using default: %vs", config.CacheTTL.Seconds())
	}
	log.Info.Printf("ProxyCacheTTL = %vs", config.CacheTTL.Seconds())

	if debugLogsEnabledStr, ok := os.LookupEnv("DEBUG"); ok {
		config.DebugLogsEnabled, _ = strconv.ParseBool(debugLogsEnabledStr)
	}
	log.Info.Printf("DebugLogsEnabled = %v", config.DebugLogsEnabled)

	return
}

// GetConfig builds the server configuration from the environment
func GetConfig() (config *ServerConfiguration, err error) {
	if cachedConfig == nil {
		cachedConfig, err = buildConfig()
	}
	return cachedConfig, err
}
