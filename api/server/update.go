package server

import (
	"api/model"
	"errors"
	"github.com/gin-gonic/gin"
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

func (h *Handlers) Move(arg model.ArgMap) ([]model.InfoList, error) {

	var infoList []model.InfoList

	fileNme := arg.SourcePathInfo[len(arg.SourcePathInfo)-1].Name
	// 需要修改的目录地址
	_, sourctPath := clearPathKey(h.WorkSpace, arg.SourcePathInfo)

	_, destionPath := clearPathKey(h.WorkSpace, arg.DestinationPathInfo)

	if exist(filepath.Join(destionPath, fileNme)) {
		return infoList, errors.New("file exist")
	}

	err := os.Rename(sourctPath, filepath.Join(destionPath, fileNme))

	return infoList, err
}

func (h *Handlers) Remove(arg model.ArgMap) ([]model.InfoList, error) {

	var infoList []model.InfoList

	// 需要删除的目录地址
	_, sourctPath := clearPathKey(h.WorkSpace, arg.PathInfo)

	err := os.RemoveAll(sourctPath)

	return infoList, err
}


func (h *Handlers) Download(c *gin.Context, arg model.ArgMap) ([]model.InfoList, error) {
	var infoList []model.InfoList

	var fileNameList []string

	for _, pathInfo := range arg.PathInfoList {
		_, pathKey := clearPathKey(h.WorkSpace, pathInfo)
		fileNameList = append(fileNameList, pathKey)
	}


	if len(fileNameList) <= 1 {
		// 文件名
		attachmentName := fileNameList[0]

		c.Header("Content-Type", "application/octet-stream")
		// c.Header("Content-Disposition", "attachment; filename="+attachmentName)
		// c.Header("Content-Transfer-Encoding", "binary")
		_, name := filepath.Split(attachmentName)
		c.FileAttachment(fileNameList[0], name)

	} else {
		attachmentName := fileNameList[0]
		path, _ := filepath.Split(attachmentName)

		// 父级文件夹名
		parentDir := filepath.Base(path)

		if len(parentDir) == 0 {
			parentDir = "mvDisk"
		}

		// 打包zip
		archNameTmp, err := createFlatZip(fileNameList, parentDir)

		if err != nil {
			return infoList, err
		}
		_, zipName := filepath.Split(archNameTmp)
		c.FileAttachment(archNameTmp, zipName)

		// 删除临时文件.zip
		_ = os.Remove(archNameTmp)
	}

	return infoList, nil
}