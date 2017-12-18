package ui

import (
	"github.com/NYTimes/gziphandler"
	"github.com/fsufitch/websnake/websnake-server/config"
	"github.com/fsufitch/websnake/websnake-server/log"
	"github.com/gorilla/mux"
)

// ApplyUIRoutes adds the UI file routes to the router
func ApplyUIRoutes(r *mux.Router) error {
	log.Debug.Print("Applying UI routes...")

	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	h := gziphandler.GzipHandler(semiStaticHandler{
		StaticDir:    config.UIStaticPath,
		CacheMap:     newCacheMap(config.CacheTTL),
		APIHostValue: config.APIHost,
	})

	r.Handle("/{filename}", h).Methods("GET")
	r.Handle("/", h).Methods("GET")

	return nil
}
