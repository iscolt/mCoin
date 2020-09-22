package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 区块
type block struct {
	data         string
	previousHash string
	hash         string
}

// 计算区块的哈希
func (b *block) initHash() {
	hashCode := computeHash(*b)
	b.hash = hashCode
}

func computeHash(b block) string {
	hash := sha256.New()
	hash.Write([]byte(b.data + b.previousHash))
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

// 区块 的 链
type chain struct {
	chain []block
}

// 生成祖先链
func bigBang() block {
	genesisBlock := block{
		data:         "祖先",
		previousHash: "0",
	}
	genesisBlock.initHash()
	return genesisBlock
}

// 获取最后一个区块
func (c chain) getLatestBlock() block {
	return c.chain[len(c.chain)-1]
}

// 添加区块到链上
func (c *chain) addBlockToChain(b block) {
	b.previousHash = c.getLatestBlock().hash
	b.initHash()
	c.chain = append(c.chain, b)
}

// 验证区块
func (c chain) validateChain() bool {
	if len(c.chain) == 1 {
		if c.chain[0].hash != computeHash(c.chain[0]) {
			return false
		}
		return true
	}

	// 从第二个区块开始验证，验证到最后一个
	for i := 1; i < len(c.chain); i++ {
		blockToValidate := c.chain[i]
		// 验证当前区块是否合法，数据是否有篡改
		if blockToValidate.hash != computeHash(blockToValidate) {
			fmt.Println("第", i, "个区块数据已被篡改")
			return false
		}
		// 验证区块的previousHash是否等于previous区块的hash
		previousBlock := c.chain[i-1]
		if previousBlock.hash != blockToValidate.previousHash {
			fmt.Println("区块", i, "与区块", i-1, "连接断裂")
			return false
		}
	}
	return true
}

func main() {
	mbChain := chain{
		chain: []block{bigBang()},
	}
	block1 := block{
		data: "test1",
	}
	mbChain.addBlockToChain(block1)
	// 篡改数据
	mbChain.chain[1].data = "test"
	fmt.Println(mbChain) //byte[]转换成string 输出
	fmt.Println(mbChain.validateChain())
}
