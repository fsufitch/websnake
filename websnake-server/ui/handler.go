package ui

import (
	"bytes"
	"errors"
	"io/ioutil"
	"mime"
	"net/http"
	"path"
	"strings"

	"github.com/gorilla/mux"
)

type semiStaticHandler struct {
	StaticDir    string
	APIHostValue string
	CacheMap     *cacheMap
}

const apiHostReplaceable = "__API_HOST__"

func (h semiStaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename, _ := mux.Vars(r)["filename"]
	if filename == "" {
		filename = "index.html"
	}

	fullPath := path.Join(h.StaticDir, filename)

	cacheEntry, ok := h.CacheMap.Get(fullPath)

	if !ok {
		data, mimeType, ok := readFile(fullPath)
		if ok {
			data = bytes.Replace(data, []byte(apiHostReplaceable), []byte(h.APIHostValue), -1)
			h.CacheMap.Set(fullPath, data, mimeType)
		} else {
			h.CacheMap.SetError(fullPath, errors.New("File not found"))
		}
		cacheEntry, _ = h.CacheMap.Get(fullPath)
	}

	if cacheEntry.Error != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(cacheEntry.Error.Error()))
		return
	}

	if r.Header.Get("If-None-Match") == cacheEntry.ETag {
		w.Header().Set("Etag", cacheEntry.ETag)
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("Content-Type", cacheEntry.ContentType)
	w.Header().Set("Etag", cacheEntry.ETag)
	w.WriteHeader(http.StatusOK)
	w.Write(cacheEntry.Data)
}

func readFile(path string) (data []byte, mimeType string, ok bool) {
	data, err := ioutil.ReadFile(path)
	ok = err == nil

	parts := strings.Split(path, ".")
	ext := "." + parts[len(parts)-1]
	mimeType = mime.TypeByExtension(ext)

	return
}
