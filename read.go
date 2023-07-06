/*
Package sherrydocument 文件管理
*/
package foldertree

import (
   "fmt"
   "os"
   "bufio"
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

// 讀取最後幾行內容
func(app *SryDocument) ReadLastNLines(fileName string, n int)([]string, error) {
   file, err := os.Open(fileName)
   if err != nil {
      return nil, err
   }
   defer file.Close()

   scanner := bufio.NewScanner(file)
   lines := make([]string, 0)

   for scanner.Scan() {
      lines = append(lines, scanner.Text())
      if len(lines) > n {
         lines = lines[1:]
      }
   }

   if scanner.Err() != nil {
      return nil, scanner.Err()
   }

   return lines, nil
}

// 讀取檔案
func(doc *SryDocument) Read(fileName string)([]byte, error) {
   s := []byte{}
   if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {  // 檔案不存在
      return  s, fmt.Errorf(fileName + " is not exist")
   }
   s, err := ioutil.ReadFile(fileName)
   if err != nil {
      return s, err
   }
   return s, nil
}

// 讀取檔案，若檔案不存在則建立空檔案
func(doc *SryDocument) ReadAndCreate(filename string)([]byte, error) {
   s := ""
   // 檔案不存在
   if _, err := os.Stat(filename); os.IsNotExist(err) {  // does not exist
      if err := doc.Create(filename, []byte(s)) ; err != nil {
         return []byte(s), nil
      }
      return []byte(s), err
   }
   return doc.Read(filename)
}
