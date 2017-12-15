package server

import (
	"fmt"
	"net/http"

	"github.com/fsufitch/websnake/websnake-server/config"
	"github.com/fsufitch/websnake/websnake-server/log"
	"github.com/fsufitch/websnake/websnake-server/ui"
	"github.com/gorilla/mux"
)

func createRoutes() (router *mux.Router, err error) {
	log.Info.Print("Creating routes...")

	router = mux.NewRouter()
	router.StrictSlash(true)

	err = ui.ApplyUIProxyRoutes(router)
	if err != nil {
		return
	}

	api := router.PathPrefix("/api").Subrouter()
	api.StrictSlash(true)
	//api.Handle("/status", status.NewHandler(apiHost, uiResURL))

	return router, nil
}

func startServer() error {
	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	router, err := createRoutes()
	if err != nil {
		return err
	}
	log.Debug.Print("Got routes:")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathSegments := []interface{}{}
		for _, a := range ancestors {
			path, _ := a.GetPathTemplate()
			pathSegments = append(pathSegments, path)
		}
		path, _ := route.GetPathTemplate()
		pathSegments = append(pathSegments, path)
		log.Debug.Print(pathSegments)
		return nil
	})

	serveAddress := fmt.Sprintf(":%d", config.Port)
	log.Info.Printf("Serving on %s", serveAddress)
	err = http.ListenAndServe(serveAddress, router)
	return err
}
