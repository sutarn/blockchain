package main

import (
	"fmt"
	"net/http"
	"github.com/winlion/blockchain/core"
)
func funcMining(w http.ResponseWriter, r *http.Request) {
	var blockchain core.BlockChain
	blockchain.Init()
	blockchain.Mining()
}
func main(){
	fmt.Println("hello,world")
	http.HandleFunc("/mining",funcMining)
	http.ListenAndServe(":80",nil)
}
