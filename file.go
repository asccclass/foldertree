/*
Package sherrydocument 文件管理
*/
package foldertree

import (
   "os"
   "errors"
   "io/ioutil"
)

// Create 建立檔案
func(doc *SryDocument) Create(path string, content []byte) (error) {
   if err := ioutil.WriteFile(path, content, 0644); err != nil {
      return err
   }
   return nil
}

// Append 擴充檔案內容
func(doc *SryDocument) Append(path string, content []byte) (error) {
   s := ""
   fx, err := os.Stat(path)
   if os.IsNotExist(err) {  // does not exist
      if err := doc.Create(path, []byte(s)) ; err != nil {
         return nil
      }
   }
   f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
   if err != nil {
      return err
   }
   defer f.Close()
   if fx.Size() != 0 {
      s = "\n"
   }
   if _, err = f.WriteString(s + string(content)); err != nil {
      return err
   }
   return nil
}

// Overwrite 覆蓋檔案，若檔案不存在則建立檔案
func(doc *SryDocument) OverWrite(path string, content []byte) (error) {
   // 檔案不存在
   if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
      return doc.Create(path, content)
   }
   // 檔案存在
   f, err := os.OpenFile(path, os.O_RDWR, 0644)
   if err != nil {
      return err
   }
   if _, err := f.Write(content); err != nil {
      f.Close() // ignore error; Write error takes precedence
      return err
   }
   if err := f.Close(); err != nil {
      return err
   }
   return nil
}

