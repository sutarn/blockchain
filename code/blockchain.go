package core

import (
	"time"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"fmt"
)

type BlockChain struct{
	Chain []Block
	Trans []Transaction
}

func (self *BlockChain)Init(){
	self.Chain = make([]Block,1)
	self.Trans = make([]Transaction,1)
}
//添加一条新的交易
func (self *BlockChain)NewTransaction(sender,recipient string ,amount Int)(lastIndex Int){
	self.Trans = append(self.Trans,Transaction{sender,recipient,amount})
	return self.LastBlock().Index
}


//最后一个索引
func (self *BlockChain)LastBlock()(Block){
	index := len(self.Chain)-1
	return self.Chain[index]
}

//添加一条新的交易
func (self *BlockChain)NewBlock(proof Int,previoushash string){
	block := Block{
		 Int(len(self.Chain) + 1),
			Int(time.Now().Unix()),
			 self.Trans,
			proof,
		previoushash,
	}
	self.Chain = append(self.Chain,block)
}



//寻找工作量参数
func (self *BlockChain)FindProof(lastProof Int) Int{
	var proof  Int = 0
	for !self.ValidProof(lastProof,proof){
		proof++
	}
	return proof
}

//工作量证明,上一个工作量参数和当前工作量参数加起来前4位为0
func (self *BlockChain)ValidProof(lastProof Int,proof Int) bool{
	h := sha256.New()
	h.Write([]byte(string(lastProof)+string(proof)))
	return strings.Index(hex.EncodeToString(h.Sum(nil)),"0000")==0
}

//采矿方法,本质上是寻找一个proof,然后添加进去
func (self *BlockChain)Mining(){
	lastBlock:=self.LastBlock();
	proof := self.FindProof(lastBlock.Proof)
	uuid := Uuid()
	self.NewTransaction("0",uuid,1)
	self.NewBlock(proof,lastBlock.PreviousHhash)
	fmt.Println(self)
}
