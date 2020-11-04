package main

import (
  ripemd1602"golang.org/x/crypto/ripemd160"
  "io"
  "os"
  "encoding/hex"
  "fmt"
)

func MyRipemd160(message string){
  hasher:=ripemd1602.New()
  hasher.Write([]byte(message))
  fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
}
