package server

import (
	"api/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)
type Handlers struct {
	WorkSpace string
	HostName  string
}

func FileManager(c *gin.Context) {
	var fileQuery model.FileManager
	var res model.ResForm
	var errorCode = 0
	res.Suc = true
	if err := c.ShouldBind(&fileQuery); err != nil {
		res.Suc = false
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

