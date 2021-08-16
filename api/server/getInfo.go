package server

import (
	"api/model"
	b64 "encoding/base64"
	"io/ioutil"
	"path/filepath"
)

/*
arg : {"pathInfo":[{"key":"Cities\\New-York","name":"New-York"}]}
 */
func(h *Handlers) GetDirContents(arg model.ArgMap) ([]model.InfoList, error) {

	var infoList []model.InfoList

	key, pathKey := clearPathKey(h.WorkSpace, arg.PathInfo)

	// 获取目录下信息
	fileInfoList, err := ioutil.ReadDir(pathKey)

	if err != nil {
		return infoList, err
	}

	for _, fileInfo := range fileInfoList {
		info := model.InfoList{}
		info.Name = fileInfo.Name()
		infoKey := key
		if len(infoKey) > 0 {
			infoKey = infoKey + "\\"
		}
		info.Key = infoKey + fileInfo.Name()
		info.DateModified = fileInfo.ModTime()
		info.IsDirectory = fileInfo.IsDir()
		info.Size = fileInfo.Size()
		info.HasSubD = hasSubDirs(pathKey, fileInfo.Name())

		if !fileInfo.IsDir() {
			fileNameBase64 := b64.URLEncoding.
				EncodeToString([]byte(filepath.Join(pathKey, fileInfo.Name())))
			info.Url = h.HostName + "/api/get-detail?filename=" + fileNameBase64
		}

		infoList = append(infoList, info)
	}

	return infoList, nil
}

// 是否有子文件夹
func hasSubDirs(p string, n string) bool {
	path := filepath.Join(p, n)
	fileInfoList, err := ioutil.ReadDir(path)

	if err != nil {
		return false
	}

	for _, fileInfo := range fileInfoList {
		if fileInfo.IsDir() {
			return true
		}
	}

	return false
}