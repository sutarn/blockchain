package core

import (
	"crypto/sha256"
	"encoding/json"
	"encoding/hex"
)

type Block struct{
	Index Int
	Timestamp Int
	Transactions []Transaction
	Proof Int
	PreviousHhash string
}

//计算一个block 的hash
func (self *Block)Hash(block Block)string{
	h := sha256.New()
	str,_:=json.Marshal(block)
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

