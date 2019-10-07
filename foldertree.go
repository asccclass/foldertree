/*
Package sherrydocument 文件管理
*/
package foldertree

import (
	"fmt"
	"os"
	"path/filepath"

	UICommunicate "github.com/asccclass/myits/libs/uicommunicate"
)

type SryFolder struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
}

type FolderTree struct {
	Name  string      `json:"name"`
	Files []SryFolder `json:"files"`
}

// SryDocument Struct 文件結構
type SryDocument struct {
	Dir  string `json:"DocumentRoot"`
	UIC  *UICommunicate.UIComm
	Tree []FolderTree `json:"foldertree"`
}

// AbsPath 轉換實際路徑
func (doc *SryDocument) AbsPath(dir string) (string, error) {
	x, err := filepath.Abs(dir) // 先轉換成絕對路徑
	if err != nil {
		return x, nil
	}
	return x, nil
}

// IsDirExist 檢查檔案目錄是否存在
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
func (doc *SryDocument) Read() {
	t := "Hello world"
	doc.UIC.ReflectMessage("WebRead", t)
}

// NewSryDocument 初始化文件管理 createdir true)若目錄不存在自動產生 false)目錄不存在返回錯誤
func NewSryDocument(uic *UICommunicate.UIComm, dir string, createdir bool) (*SryDocument, error) {
	if dir == "" {
		return nil, fmt.Errorf("no such dir:%s", dir)
	}

	doc := &SryDocument{
		Dir: dir,
		UIC: uic,
	}

	dir, err := doc.AbsPath(dir) // 轉換絕對路徑
	if err != nil {
		return nil, err
	}

	if err := doc.IsDirExist(dir, createdir); err != nil { // 檢查目錄是否存在
		return nil, err
	}

	doc.Dir = dir
	return doc, nil
}
