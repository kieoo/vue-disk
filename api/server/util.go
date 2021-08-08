package server

import (
	"api/model"
	"path/filepath"
	"strings"
)

func clearPathKey(workspace string, pathInfo []model.PathMap) (key string, pathKey string) {
	if len(pathInfo) <= 0 {
		pathInfo = append(pathInfo, model.PathMap{Key:"", Name:""})
	}
	for _, path := range pathInfo {
		pathKey = filepath.Join(workspace)
		key = path.Key
		sonPaths := strings.Split(path.Key, "\\")
		// 目录拼接
		for _, son := range sonPaths {
			pathKey = filepath.Join(pathKey, son)
		}
	}

	return key, pathKey
}
