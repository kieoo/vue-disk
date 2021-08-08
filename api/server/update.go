package server

import (
	"api/model"
	"os"
	"path/filepath"
)

func (h *Handlers) Rename(arg model.ArgMap) ([]model.InfoList, error) {
	var infoList []model.InfoList

	// 需要修改的目录地址
	_, pathKey := clearPathKey(h.WorkSpace, arg.PathInfo)

	parentPathInfo := arg.PathInfo[0:len(arg.PathInfo)-1]

	// 父级地址
	_, parentPathKey := clearPathKey(h.WorkSpace, parentPathInfo)

	// 重命名
	err := os.Rename(pathKey, filepath.Join(parentPathKey,arg.Name))

	if err != nil {
		return infoList, err
	}
	return infoList, nil
}
