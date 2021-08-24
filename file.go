/*
Package sherrydocument 文件管理
*/
package foldertree

import (
   "io/ioutil"
   "os"
)

// Create 建立檔案
func(doc *SryDocument) Create(path string, content []byte) (error) {
   if err := ioutil.WriteFile(path, content, 0644); err != nil {
      return err
   }
   return nil
}

func(doc *SryDocument) Append(path string, content []byte) (error) {
   if _, err := os.Stat(path); os.IsNotExist(err) {  // does not exist
      s := ""
      if err := doc.Create(path, []byte(s)) ; err != nil {
         return nil
      }
   }
   f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
   if err != nil {
      return err
   }
   defer f.Close()
   if _, err = f.WriteString("\n" + string(content)); err != nil {
      return err
   }
   return nil
}
