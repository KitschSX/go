package main

import (
  "bytes"
  "crypto/sha256"
  "encoding/binary"
  "encoding/json"
  "fmt"
  "log"
  "math"
  "math/big"
  "net/http"
  "strconv"
  "time"
)

// Nonce 在对工作量证明进行验证时用到
type Block struct {
  Timestamp     int64
  PrevBlockHash []byte
  Hash          []byte
  Data          []byte
  Nonce         int
}
type blo struct {
  PrevBlockHash string
  Hash          string
  Data          string
}

// 创建新块时需要运行工作量证明找到有效哈希
func NewBlock(data string, prevBlockHash []byte) *Block {
  block := &Block{
    Timestamp:     time.Now().Unix(),
    PrevBlockHash: prevBlockHash,
    Hash:          []byte{},
    Data:          []byte(data),
    Nonce:         0}
  pow := NewProofOfWork(block)
  nonce, hash := pow.Run()

  block.Hash = hash[:]
  block.Nonce = nonce

  return block
}

func NewGenesisBlock() *Block {
  return NewBlock("Genesis Block", []byte{})
}
type BlockChain struct {
  blocks []*Block
}

func NewBlockChain() *BlockChain {
  return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func (bc *BlockChain) AddBlock(data string) {
  prevBlock := bc.blocks[len(bc.blocks)-1]
  newBlock := NewBlock(data, prevBlock.Hash)
  bc.blocks = append(bc.blocks, newBlock)
}
func main() {
  bc := NewBlockChain()

  bc.AddBlock("pay 1 BTC to xiaoming")
  bc.AddBlock("send 2  BTC to kk")

  for _, block := range bc.blocks {

    fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
    fmt.Printf("Data: %s\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)
    pow := NewProofOfWork(block)
    fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
  }
  var ss []blo
  http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
   //跨域问题
   //writer.Header().Set("Access-Control-Allow-Origin", "*")
   //获取url中的数据
   data := request.FormValue("data")
   bc.AddBlock(data)
   var bloc blo
   ss = append(ss,bloc)

   for i, block := range bc.blocks {
      fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
      fmt.Printf("Data: %s\n", block.Data)
      fmt.Printf("Hash: %x\n", block.Hash)
      pow := NewProofOfWork(block)
      fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))

      ss[i].Data=string(block.Data)
      ss[i].PrevBlockHash=string(block.PrevBlockHash)
      ss[i].Hash=string(block.Hash)
   }


    bytes, _ := json.Marshal(ss)
    fmt.Fprint(writer,string(bytes))
  })
  fmt.Println(http.ListenAndServe(":8080",nil))
}
// 难度值，这里表示哈希的前 12 位必须是 0
const targetBits = 12

const maxNonce = math.MaxInt64

// 每个块的工作量都必须要证明，所有有个指向 Block 的指针
// target 是目标，我们最终要找的哈希必须要小于目标
type ProofOfWork struct {
  block  *Block
  target *big.Int
}

// target 等于 1 左移 256 - targetBits 位
func NewProofOfWork(b *Block) *ProofOfWork {
  target := big.NewInt(1)
  target.Lsh(target, uint(256-targetBits))

  pow := &ProofOfWork{b, target}

  return pow
}

// 工作量证明用到的数据有：PrevBlockHash, Data, Timestamp, targetBits, nonce
func (pow *ProofOfWork) prepareData(nonce int) []byte {
  data := bytes.Join(
    [][]byte{
      pow.block.PrevBlockHash,
      pow.block.Data,
      IntToHex(pow.block.Timestamp),
      IntToHex(int64(targetBits)),
      IntToHex(int64(nonce)),
    },
    []byte{},
  )

  return data
}

// 工作量证明的核心就是寻找有效哈希
func (pow *ProofOfWork) Run() (int, []byte) {
  var hashInt big.Int
  var hash [32]byte
  nonce := 0

  fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
  for nonce < maxNonce {
    data := pow.prepareData(nonce)

    hash = sha256.Sum256(data)
    hashInt.SetBytes(hash[:])

    if hashInt.Cmp(pow.target) == -1 {
      fmt.Printf("\r%x", hash)
      break
    } else {
      nonce++
    }
  }
  fmt.Print("\n\n")

  return nonce, hash[:]
}

// 验证工作量，只要哈希小于目标就是有效工作量
func (pow *ProofOfWork) Validate() bool {
  var hashInt big.Int

  data := pow.prepareData(pow.block.Nonce)
  hash := sha256.Sum256(data)
  hashInt.SetBytes(hash[:])

  isValid := hashInt.Cmp(pow.target) == -1

  return isValid
}

// 将一个 int64 转化为一个字节数组(byte array)
func IntToHex(num int64) []byte {
  buff := new(bytes.Buffer)
  err := binary.Write(buff, binary.BigEndian, num)
  if err != nil {
    log.Panic(err)
  }

  return buff.Bytes()
}
