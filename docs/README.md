## FolderTree 目錄檔案管理工具
主要用來管理特定目錄下之目錄及檔案，方便其他程式使用。

### INstall
```
go get github.com/asccclass/foldertree
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
	trees, err := foldertree.NewSryDocument("./foldertree", false)
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