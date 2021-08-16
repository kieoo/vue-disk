package server

import (
	"api/model"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// arguments : {"pathInfo":[{"key":"Landscapes","name":"Landscapes"}],"name":"test"}

func (h *Handlers) CreateDir(arg model.ArgMap) ([]model.InfoList, error) {

	var infoList []model.InfoList

	_, pathKey := clearPathKey(h.WorkSpace, arg.PathInfo)

	directory := filepath.Join(pathKey, arg.Name)

	if exist(directory) {
		return infoList, errors.New("directory exit")
	}

	// 创建
	err := os.Mkdir(directory, os.ModePerm)

	if err != nil {
		return infoList, err
	}
	return infoList, nil
}

//arguments: {"destinationPathInfo":[{"key":"test","name":"test"},{"key":"test\\\\test_create","name":"test_create"}],"chunkMetadata":"{\"UploadId\":\"12c4c6f0-733b-2ed6-d134-67aa57c68c4c\",\"FileName\":\"kingsoft-wpsmail1-slow.log\",\"Index\":0,\"TotalCount\":3,\"FileSize\":598087}"}

func (h * Handlers) Upload(c *gin.Context ,arg model.ArgMap) ([]model.InfoList, error) {

	var infoList []model.InfoList
	_, pathKey := clearPathKey(h.WorkSpace, arg.DestinationPathInfo)

	file, _, _ := c.Request.FormFile("chunk")

	// 保存
	var chunkMap model.ChunkMetadataMap
	_ = json.Unmarshal([]byte(arg.ChunkMetadata), &chunkMap)
	fileName := filepath.Join(pathKey, chunkMap.FileName)

	err := fileChunkWrite(file, fileName, chunkMap)

	if err != nil {
		return infoList, err
	}

	return infoList, nil
}

func (h *Handlers)AbortUpload(arg model.ArgMap) ([]model.InfoList, error) {
	var infoList []model.InfoList
	uploadingFile, ok := Uploading[arg.UploadId]
	if ok {
		go func(f string) {
			if !exist(uploadingFile) {
				time.Sleep(5 * time.Second)
			}
			_ = os.RemoveAll(uploadingFile)
		}(uploadingFile)

		delete(Uploading, arg.UploadId)
	}
	return infoList, nil
}

//{"sourcePathInfo":[{"key":"test","name":"test"},{"key":"test\\test1_1","name":"test1_1"}],"sourceIsDirectory":true,"destinationPathInfo":[{"key":"test2al","name":"test2al"}]}
func (h *Handlers) Copy(arg model.ArgMap) ([]model.InfoList, error) {

	var infoList []model.InfoList

	fileNme := arg.SourcePathInfo[len(arg.SourcePathInfo)-1].Name
	// 需要修改的目录地址
	_, sourctPath := clearPathKey(h.WorkSpace, arg.SourcePathInfo)

	_, destionPath := clearPathKey(h.WorkSpace, arg.DestinationPathInfo)

	if exist(filepath.Join(destionPath, fileNme)) {
		return infoList, errors.New("file exist")
	}

	err := deepCopy(sourctPath, destionPath)

	return infoList, err
}

func fileChunkWrite(file multipart.File, path string, cm model.ChunkMetadataMap) error {

	if exist(path) {
		return errors.New("files exit")
	}

	buf := make([]byte, 0)

	filePath := path + "." + cm.UploadId

	Uploading[cm.UploadId] = filePath

	buf, _ = ioutil.ReadAll(file)

	fd, _ := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)

	if _, err := fd.Write(buf); err != nil {
		return err
	}

	fd.Close()

	if cm.Index +1 >= cm.TotalCount {
		os.Rename(filePath, path)
		delete(Uploading, cm.UploadId)
	}
	return nil
}




