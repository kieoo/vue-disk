package server

import (
	"api/model"
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
			info.Url = h.HostName + "/getDetail?filename=info.Key"
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