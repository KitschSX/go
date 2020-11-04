package main

import (
  "crypto"
  "crypto/md5"
  "crypto/rand"
  "crypto/rsa"
  "fmt"
)

func main() {
  //生成私钥
  priv,_:=rsa.GenerateKey(rand.Reader,1024)
  pub:=&priv.PublicKey
  plaitxt:=[]byte("hello world")
  h:=md5.New()
  h.Write(plaitxt)
  hashed:=h.Sum(nil)
  opts:=rsa.PSSOptions{rsa.PSSSaltLengthAuto,crypto.MD5}
  sig,_:=rsa.SignPSS(rand.Reader,priv,crypto.MD5,hashed,&opts)
  fmt.Println(sig)
  err:=rsa.VerifyPSS(pub,crypto.MD5,hashed,sig,&opts)
  if err==nil {
    fmt.Println("验签成功")
  }
}
