/*
Package sherrydocument 文件管理
*/
package foldertree

import (
   "os"
   "errors"
   "io/ioutil"
   // "net/http"
   // "github.com/gorilla/mux"
)
/*
//Read file
func(app *SryDocument) ReadfileFromWeb(w http.ResponseWriter, r *http.Request) {
   webVars := mux.Vars(r)
   if webVars["fileName"] == "" {
      app.Srv.Error.Error2Web(w, fmt.Errorf("file name is empty"))
      return
   }
   s, err := os.ReadFile(app.PoolPath + webVars["fileName"])
   if err != nil {
      app.Srv.Error.Error2Web(w, fmt.Errorf("file name is empty"))
      return
   }
   mime := http.DetectContentType(s)
   fileSize := len(string(s))

   w.Header().Set("Content-Type", mime)
   w.Header().Set("Content-Disposition", "attachment; filename=" + webVars["fileName"])
   w.Header().Set("Content-Length", strconv.Itoa(fileSize))
   http.ServeContent(w, r, webVars["fileName"], time.Now(), bytes.NewReader(s))
}
*/

// 讀取檔案
func(doc *SryDocument) Read(fileName string)([]byte, (error) {
   s := []byte
   if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {  // 檔案不存在
      return  s, fmt.Errorf(fileName + " is not exist")
   }
   s, err := ioutil.ReadFile(fileName)
   if err != nil {
      return s, err
   }
   return s, nil
}

