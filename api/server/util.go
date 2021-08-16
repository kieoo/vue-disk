package server

import (
	"api/model"
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
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


func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func createFlatZip(filesList []string, archName string) (string, error) {

	t := time.Now().Unix()

	tmpPath, _ := os.Getwd()

	tmpPath = filepath.Join(tmpPath, "tmp")

	_ = os.MkdirAll(tmpPath, os.ModePerm)

	archNameTmp := filepath.Join(tmpPath, fmt.Sprintf("%s_%d.zip", archName, t))

	// 先删除临时文件
	_ = os.RemoveAll(archNameTmp)

	zf, err := os.OpenFile(archNameTmp, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return archNameTmp, err
	}
	defer zf.Close()

	w := zip.NewWriter(zf)
	defer w.Close()

	for _, file := range filesList {
		src, err := os.Open(file)
		if err != nil {
			return archNameTmp, err
		}

		_, filename := filepath.Split(file)
		f, err := w.Create(filename)

		if err != nil {
			_ = src.Close()
			return archNameTmp, err
		}

		_, err = io.Copy(f, src)

		if err != nil {
			_ = src.Close()
			return archNameTmp, err
		}

		_ = src.Close()
	}

	return archNameTmp, nil
}

// src: 原地址
// dis: 目标目录
func deepCopy(src string, dist string) error {
	f, err := os.Stat(src)
	if err != nil {
		return err
	}
	if f.IsDir() {
		fileInfos, err := ioutil.ReadDir(src)

		if err != nil {
			return err
		}

		err = os.MkdirAll(filepath.Join(dist, f.Name()), os.ModePerm)
		if err != nil {
			return err
		}

		dist = filepath.Join(dist, f.Name())

		for _, fileInfo := range fileInfos {
			 err = deepCopy(filepath.Join(src, fileInfo.Name()), dist)
             if err != nil {
             	return err
			 }
		}
	} else {
		srcFile, _ := os.Open(src)
		defer srcFile.Close()

		_, fileName := filepath.Split(src)
		disFile, _ := os.OpenFile(filepath.Join(dist, fileName), os.O_CREATE|os.O_RDWR, os.ModePerm)
		defer disFile.Close()

		_, err := io.Copy(disFile, srcFile)
		if err != nil {
			return err
		}
	}

	return nil
}