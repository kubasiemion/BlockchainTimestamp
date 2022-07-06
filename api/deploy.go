package api

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func DeployCertContract(chainid int) (common.Address, *types.Transaction, *Api, error) {

	client, err := Dial(chainid)
	if err != nil {
		panic(err)
	}

	publicKey := GetEthKey().PublicKey //privateKey.Public()

	fromAddress := crypto.PubkeyToAddress(publicKey)
	nonce, err := GetNonceForAddress(client, fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(GetEthKey(), chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units

	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	//auth.GasPrice = big.NewInt(35680970134)

	return DeployApi(auth, client)

}
