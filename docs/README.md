## FolderTree 目錄檔案管理工具
主要用來管理特定目錄下之目錄及檔案，方便其他程式使用。

### Install
```
go get github.com/asccclass/foldertree
```

* Initial

```
// 初始化文件管理 createdir true)若目錄不存在自動產生 false)目錄不存在返回錯誤
NewSryDocument(system, dir string, createdir bool) (*SryDocument, error)

// 參數說明：
   System: 系統 windows or Linux or Mac
```

### 可用函數
* Create(path, content []byte)(error)			// 建立檔案
* Append(path string, content []byte) (error)		// 擴充檔案內容，若檔案不存在則建立該檔案
* OverWrite(path string, content []byte) (error)	// 覆蓋檔案，若檔案不存在會建立檔案
* Read(fileName string)([]byte, error)			// 讀取檔案
* ReadAndCreate(filename string)([]byte, error 		// 取檔案，若檔案不存在則建立空檔案  
* ReadLastNLines(fileName string, n int) ([]string, error) // 讀取最後幾行內容
* IsDirExist(dir string, created bool)(error)		// 判斷 works 目錄是否存在，若不存在則建立新目錄 
* AbsPath(dir string) (string, error)			// 轉換實際（絕對）路徑 

### 結構 struct

```
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
```

### Usage
```
package main

import (
   "encoding/json"
   "fmt"

   "github.com/asccclass/foldertree"
)

func main() {
   // (*SryDocument, error)
   trees, err := foldertree.NewSryDocument("windows", "./foldertree", false)
   if err != nil {
      fmt.Println(err)
   }
   b, err := json.Marshal(&trees)
   if err != nil {
      fmt.Println(err)
   }
   fmt.Println(string(b))

}
```
