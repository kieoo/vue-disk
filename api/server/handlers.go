package server

import (
	"api/constant"
	"api/model"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

// 记录正在上传的文件
var Uploading = make(map[string]string)

type Handlers struct {
	WorkSpace string
	HostName  string
}

func FileManager(c *gin.Context) {
	var fileQuery model.FileManager
	var res model.ResForm
	var errorCode = constant.Other
	res.Suc = true
	if err := c.ShouldBind(&fileQuery); err != nil {
		res.Suc = false
		errorCode = constant.DirectoryNotFound
		res.ErrorCode = &errorCode
		res.ErrorText = fmt.Sprintf("%s", err)
		c.JSONP(http.StatusBadRequest, res)
		return
	}

	pwd,_ := os.Getwd()
	pwd = filepath.Join(pwd, "mdisk")

	// 不存在创建目录
	err := os.MkdirAll(pwd, os.ModePerm)
	if  err != nil {
		res.Suc = false
		errorCode = constant.DirectoryNotFound
		res.ErrorCode = &errorCode
		res.ErrorText = fmt.Sprintf("%s", err)
		c.JSONP(http.StatusBadRequest, res)
		return
	}

	host :=  "http://" + c.Request.Host
	handler := Handlers{pwd, host}

	switch fileQuery.Com {
	case "GetDirContents":
		result, err := handler.GetDirContents(fileQuery.Arg)
		if err != nil {
			res.Suc = false
			errorCode = constant.DirectoryNotFound
			res.ErrorCode = &errorCode
			res.ErrorText = fmt.Sprintf("%s", err)
			c.JSONP(http.StatusBadRequest, res)
			return
		}

		res.Suc = true
		res.Result = &result
		c.JSONP(http.StatusOK, res)
		return

	case "CreateDir":
		_, err := handler.CreateDir(fileQuery.Arg)
		if err != nil {
			res.Suc = false
			errorCode = constant.DirectoryExists
			res.ErrorCode = &errorCode
			res.ErrorText = fmt.Sprintf("%s", err)
			c.JSONP(http.StatusBadRequest, res)
			return
		}
		res.Suc = true
		// res.Result = &result
		c.JSONP(http.StatusOK, res)

	case "UploadChunk":
		_, err := handler.Upload(c, fileQuery.Arg)
		if err != nil {
			res.Suc = false
			errorCode = constant.FileExists
			res.ErrorCode = &errorCode
			res.ErrorText = fmt.Sprintf("%s", err)
			c.JSONP(http.StatusBadRequest, res)
			return
		}
		res.Suc = true
		// res.Result = &result
		c.JSONP(http.StatusOK, res)

	case "AbortUpload":
		_, err := handler.AbortUpload(fileQuery.Arg)
		if err != nil {
			res.Suc = false
			errorCode = constant.WrongFileExtension
			res.ErrorCode = &errorCode
			res.ErrorText = fmt.Sprintf("%s", err)
			c.JSONP(http.StatusBadRequest, res)
			return
		}
		res.Suc = true
		// res.Result = &result
		c.JSONP(http.StatusOK, res)

	case "Rename":
		_, err := handler.Rename(fileQuery.Arg)
		if err != nil {
			res.Suc = false
			errorCode = constant.FileNotFound
			res.ErrorCode = &errorCode
			res.ErrorText = fmt.Sprintf("%s", err)
			c.JSONP(http.StatusBadRequest, res)
			return
		}
		res.Suc = true
		// res.Result = &result
		c.JSONP(http.StatusOK, res)

	case "Download":
		_, err := handler.Download(c, fileQuery.Arg)
		if err != nil {
			res.Suc = false
			errorCode = constant.FileNotFound
			res.ErrorCode = &errorCode
			res.ErrorText = fmt.Sprintf("%s", err)
			c.JSONP(http.StatusBadRequest, res)
			return
		}
		res.Suc = true
		// res.Result = &result
		// c.JSONP(http.StatusOK, res)

	case "Move":
		_, err := handler.Move(fileQuery.Arg)
		if err != nil {
			res.Suc = false
			errorCode = constant.WrongFileExtension
			res.ErrorCode = &errorCode
			res.ErrorText = fmt.Sprintf("%s", err)
			c.JSONP(http.StatusBadRequest, res)
			return
		}
		res.Suc = true
		// res.Result = &result
		c.JSONP(http.StatusOK, res)

	case "Remove":
		_, err := handler.Remove(fileQuery.Arg)
		if err != nil {
			res.Suc = false
			errorCode = constant.FileNotFound
			res.ErrorCode = &errorCode
			res.ErrorText = fmt.Sprintf("%s", err)
			c.JSONP(http.StatusBadRequest, res)
			return
		}
		res.Suc = true
		// res.Result = &result
		c.JSONP(http.StatusOK, res)

	case "Copy":
		_, err := handler.Copy(fileQuery.Arg)
		if err != nil {
			res.Suc = false
			errorCode = constant.FileExists
			res.ErrorCode = &errorCode
			res.ErrorText = fmt.Sprintf("%s", err)
			c.JSONP(http.StatusBadRequest, res)
			return
		}
		res.Suc = true
		// res.Result = &result
		c.JSONP(http.StatusOK, res)

	}
}

func GetDetail(c *gin.Context) {

	filenameb64, _ := c.GetQuery("filename")
	var res model.ResForm

	if len(filenameb64) <= 0 {
		res.Suc = false
		errorCode := constant.FileExists
		res.ErrorCode = &errorCode
		res.ErrorText = fmt.Sprintf("%s", "file not exist")
		c.JSONP(http.StatusBadRequest, res)
		return
	}

	fileName, _ := base64.URLEncoding.DecodeString(filenameb64)
	c.File(string(fileName))
}

