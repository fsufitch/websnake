package ui

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/fsufitch/websnake/websnake-server/config"
	"github.com/fsufitch/websnake/websnake-server/log"
	"github.com/fsufitch/websnake/websnake-server/util"
	"github.com/gorilla/mux"
)

// ProxyHandler handles displaying the basics of the user interface
type ProxyHandler struct {
	apiHost   string
	targetURL *url.URL
	cache     *proxyCacheMap
}

// NewProxyHandler creates a new proxy handler for a given URL target and TTL
func NewProxyHandler(apiHost string, target *url.URL, cacheTTL time.Duration) *ProxyHandler {
	return &ProxyHandler{
		apiHost:   apiHost,
		targetURL: target,
		cache:     newProxyCacheMap(cacheTTL),
	}
}

func (h ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename, _ := mux.Vars(r)["filename"]

	if filename == "" {
		filename = "index.html"
	}

	fileURL, err := url.Parse(filename)
	if err != nil {
		util.WriteHTTPErrorResponse(w, 500, err.Error())
		return
	}

	cacheEntry, cacheHit := h.cache.Get(filename)
	if !cacheHit {
		if data, contentType, err := h.fetchProxiedFile(fileURL); err == nil {
			cacheEntry = h.cache.Set(filename, data, contentType)
		} else {
			cacheEntry = h.cache.SetError(filename, err)
		}
	}

	if cacheEntry.Error == nil {
		w.Header().Set("Content-Type", cacheEntry.ContentType)
		w.WriteHeader(200)
		w.Write(cacheEntry.Data)
		return
	}

	if cacheEntry.Error == errProxiedFileNotFound {
		util.WriteHTTPErrorResponse(w, 404, "File not found")
		return
	}

	util.WriteHTTPErrorResponse(w, 502, cacheEntry.Error.Error())
}

var errProxiedFileNotFound = errors.New("File not found")

func (h ProxyHandler) fetchProxiedFile(fileURL *url.URL) (data []byte, contentType string, err error) {
	url := h.targetURL.ResolveReference(fileURL)
	var response *http.Response
	done := make(chan bool)

	log.Info.Printf("Fetching proxied file: %s", url.String())
	go func() {
		response, err = http.Get(url.String())
		done <- true
	}()

	<-done

	if err != nil {
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		data, _ = ioutil.ReadAll(response.Body)
		contentType = response.Header.Get("Content-Type")

		data = []byte(strings.Replace(string(data), "__API_HOST__", h.apiHost, -1))
	case http.StatusNotFound:
		err = errProxiedFileNotFound
	default:
		errData, _ := ioutil.ReadAll(response.Body)
		err = fmt.Errorf("Proxy failure, status %d, content: %s", response.StatusCode, string(errData))
	}

	return
}

// ApplyUIProxyRoutes adds the UI file proxy routes to the router
func ApplyUIProxyRoutes(r *mux.Router) error {
	log.Debug.Print("Applying UI proxy routes...")

	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	h := gziphandler.GzipHandler(NewProxyHandler(
		config.APIHost,
		config.BaseUIResourcesURL,
		config.ProxyCacheTTL,
	))

	r.Handle("/{filename}", h).Methods("GET")
	r.Handle("/", h).Methods("GET")

	return nil
}
