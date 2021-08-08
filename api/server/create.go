package server

import (
	"api/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

// arguments : {"pathInfo":[{"key":"Landscapes","name":"Landscapes"}],"name":"test"}

func (h *Handlers) CreateDir(arg model.ArgMap) ([]model.InfoList, error) {
	var infoList []model.InfoList

	_, pathKey := clearPathKey(h.WorkSpace, arg.PathInfo)

	directory := filepath.Join(pathKey, arg.Name)

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

func fileChunkWrite(file multipart.File, path string, cm model.ChunkMetadataMap) error {
	var Buf = make([]byte, 0)
	Buf, _ = ioutil.ReadAll(file)
	filePath := path + "." + cm.UploadId
	fd, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if _, err := fd.Write(Buf); err != nil {
		fd.Close()
		return err
	}
	fd.Close()

	if cm.Index +1 >= cm.TotalCount {
		os.Rename(filePath, path)
	}
	return nil
}




