/*
Package sherrydocument 文件管理
*/
package foldertree

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type FolderTree struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	// Mode    os.FileMode  `json:mode"`
	ModTime time.Time    `json:"time"`
	IsDir   bool         `json:"isdir"` // 是否為目錄
	Trees   []FolderTree `json:"subnode"`
}

// SryDocument Struct 文件結構
type SryDocument struct {
	System string       `json:"system"`       // 系統 windows or Linux or Mac
	Dir    string       `json:"DocumentRoot"` // 根目錄
	Trees  []FolderTree `json:"foldertree"`
}

// AbsPath 轉換實際路徑
func (doc *SryDocument) AbsPath(dir string) (string, error) {
	x, err := filepath.Abs(dir) // 先轉換成絕對路徑
	if err != nil {
		return x, nil
	}
	return x, nil
}

// IsDirExist 檢查檔案目錄是否存在＆建立目錄
func (doc *SryDocument) IsDirExist(dir string, created bool) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) { // 目錄不存在
		if !created {
			return err
		}
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

// Read 讀取目錄資訊
func (doc *SryDocument) ParseTree(path string) ([]FolderTree, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	del := "/"
	if doc.System == "windows" {
		del = `\`
	}
	var trees []FolderTree
	for _, file := range files {
		f := &FolderTree{
			Name:  file.Name(),
			IsDir: file.IsDir(),
		}
		if file.IsDir() {
			f.Trees, err = doc.ParseTree(path + del + file.Name())
			if err != nil {
				return nil, err
			}
		} else {
			f.Size = file.Size()
			f.ModTime = file.ModTime()
		}
		trees = append(trees, *f)
	}
	return trees, nil
}

// Create 建立檔案
func (doc *SryDocument) Create(path, content string) error {
	if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}
	return nil
}

// NewSryDocument 初始化文件管理 createdir true)若目錄不存在自動產生 false)目錄不存在返回錯誤
func NewSryDocument(system, dir string, createdir bool) (*SryDocument, error) {
	if dir == "" {
		return nil, fmt.Errorf("no such dir:%s", dir)
	}

	doc := &SryDocument{
		System: system,
		Dir:    dir,
	}

	dir, err := doc.AbsPath(dir) // 轉換絕對路徑
	if err != nil {
		return nil, err
	}

	if err := doc.IsDirExist(dir, createdir); err != nil { // 檢查目錄是否存在
		return nil, err
	}

	doc.Dir = dir
	doc.Trees, err = doc.ParseTree(dir) // parse directory from root dir
	if err != nil {
		return nil, nil
	}
	return doc, nil
}
