package main

import (
  "crypto/md5"
  "crypto/rand"
  "crypto/rsa"
  "encoding/base64"
  "fmt"
)

func main() {
  priv,_:=rsa.GenerateKey(rand.Reader,1024)
  fmt.Println("私钥为：",priv)
  //通过私钥创建公钥
  pub:=priv.PublicKey//加密和解密
  org:=[]byte("hello China")//oaep实现公钥加密
  cipherTxt,_:=rsa.EncryptOAEP(md5.New(),rand.Reader,&pub,org,nil)//打印密文
  fmt.Println(cipherTxt)
  fmt.Println(base64.StdEncoding.EncodeToString(cipherTxt))
  //解密
  plaintext,_:=rsa.DecryptOAEP(md5.New(),rand.Reader,priv,cipherTxt,nil)//打印明文
  fmt.Println(plaintext)
  fmt.Println(string(plaintext))
}
