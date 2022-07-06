package main

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/kubasiemion/blockchaintimestamp/api"
)

func main() {

	n, err := api.FindCertBySerial(big.NewInt(1234567890), 80001)
	fmt.Println(err)
	b, err := json.MarshalIndent(n, " ", " ")
	fmt.Println(err)
	fmt.Println(string(b))
	fmt.Println(n.BlockTimeString())

	n, err = api.FindCertBySerial(big.NewInt(1234567890), 11155111)
	fmt.Println(err)
	b, err = json.MarshalIndent(n, " ", " ")
	fmt.Println(err)
	fmt.Println(string(b))
	fmt.Println(n.BlockTimeString())

	/*
		hash := sha256.Sum256([]byte("Dupa Jasio Karuzela"))
		tx, err := api.Mint(80001, common.HexToAddress("0x7955Dd0b14b234CAecb5974B929935f16675Cc47"), big.NewInt(1234567890), hash)
		fmt.Println(err)
		fmt.Println(tx.Hash().Hex())
		tx, err = api.Mint(11155111, common.HexToAddress("0x7955Dd0b14b234CAecb5974B929935f16675Cc47"), big.NewInt(1234567890), hash)
		fmt.Println(err)
		fmt.Println(tx.Hash().Hex())
	*/
	//fmt.Println(api.DeployCertContract(80001))
}

// Matic Cert 0x67cf9af1746d36e0d08947ff4beb167ee1da39c1
//0xf439f261bdfdeb3204938d51e131dc988497d03c
// Sepolia Cert 0x8596e9c3adfe513198c0462a1f2511fcb4cf97a7

//Sepolia2 0x9e5cdfb8c8498d0510e1ca7b1d224eca2a39259d
//Mumbai2
