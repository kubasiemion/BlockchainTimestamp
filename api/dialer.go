package api

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	kmsclitool "github.com/proveniencenft/kmsclitool/common"
)

var apibychainid = map[int][]string{ // chainid -> [name, url]
	80001:    {"Polygon Testnet Mumbai", "https://matic-mumbai.chainstacklabs.com"},
	1:        {"Ethereum Mainnet"},
	3:        {"Ethereum Testnet Ropsten", "https://ropsten.infura.io/v3/9580584257984a17927f94f2dd44aa46"},
	4:        {"Ethereum Testnet Rinkeby"},
	5:        {"Ethereum Testnet Goerli", "https://goerli.infura.io/v3/9580584257984a17927f94f2dd44aa46"},
	11155111: {"Ethereum Sepolia", "https://rpc.sepolia.dev/"}, //"https://nunki.htznr.fault.dev/rpc"},

	42: {"Ethereum Testnet Kovan"},
}

// Matic Cert 0x67cf9af1746d36e0d08947ff4beb167ee1da39c1
// Sepolia Cert 0x8596e9c3adfe513198c0462a1f2511fcb4cf97a7
var DeployedByChain = map[int]common.Address{
	11155111: common.HexToAddress("0x2834d73665d4d1be0f1e703a14a69a4588c867d9"),
	80001:    common.HexToAddress("0xf439f261bdfdeb3204938d51e131dc988497d03c"),
}

var ethkey *ecdsa.PrivateKey

func init() {

	kf, e := kmsclitool.ReadKeyfile("test746d2391a33.json")
	if e != nil {
		panic(e)
	}
	e = kmsclitool.DecryptKeyFile(kf, "kaczuszka")
	if e != nil {
		panic(e)
	}
	hextxt := hex.EncodeToString(kf.Plaintext)
	ethkey, e = crypto.HexToECDSA(hextxt)
	if e != nil {
		panic(e)
	}

}

func GetEthKey() *ecdsa.PrivateKey {
	return ethkey
}

func Dial(chainid int) (*ethclient.Client, error) {
	//see in the cache
	if client, ok := clients[chainid]; ok {
		return client, nil
	}

	nets, ok := apibychainid[chainid]
	if !ok {
		return nil, fmt.Errorf("Unknown chain")
	}
	if len(nets) < 2 {
		return nil, fmt.Errorf("Unknown API endpoint for %v", chainid)
	}
	var err error
	apiaddr := nets[1]
	client, err := ethclient.Dial(apiaddr)
	if err != nil {
		clients[chainid] = client
	}
	return client, err

}

var clients = map[int]*ethclient.Client{}

func GetNonceForAddress(client *ethclient.Client, addr common.Address) (uint64, error) {

	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func getNonce(client *ethclient.Client) (uint64, error) {

	addr := crypto.PubkeyToAddress(ethkey.PublicKey)
	return GetNonceForAddress(client, addr)
}

func GetChainID(client *ethclient.Client) (*big.Int, error) {
	return client.ChainID(context.Background())
}

func FindCertByIssue(issue, chainid int) (*Note, error) {
	cotractAddress, ok := DeployedByChain[chainid]
	if !ok {
		return nil, fmt.Errorf("Contract not deployed on chain %v", chainid)
	}

	client, err := Dial(chainid)
	if err != nil {
		return nil, err
	}

	CertAPI, err := NewApi(cotractAddress, client)
	if err != nil {
		fmt.Println("Error getting API")
		panic(err)
	}

	return CertAPI.CertsByIssue(&bind.CallOpts{}, uint8(issue))

}

func FindCertBySerial(serial *big.Int, chainid int) (*Note, error) {
	cotractAddress, ok := DeployedByChain[chainid]
	if !ok {
		return nil, fmt.Errorf("Contract not deployed on chain %v", chainid)
	}

	client, err := Dial(chainid)
	if err != nil {
		return nil, err
	}

	CertAPI, err := NewApi(cotractAddress, client)
	if err != nil {
		fmt.Println("Error getting API")
		panic(err)
	}

	return CertAPI.AllCerts(&bind.CallOpts{}, serial)

}

func Mint(chainid int, holder common.Address, serial *big.Int, hash [32]byte) (*types.Transaction, error) {

	cotractAddress, ok := DeployedByChain[chainid]
	if !ok {
		return nil, fmt.Errorf("Contract not deployed on chain %v", chainid)
	}

	client, err := Dial(chainid)
	if err != nil {
		panic(err)
	}

	privateKey := GetEthKey()
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := GetNonceForAddress(client, fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5000000) // in units

	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	//auth.GasPrice = big.NewInt(35680970134)
	CertAPI, err := NewApi(cotractAddress, client)
	if err != nil {
		fmt.Println("Error getting API")
		panic(err)
	}
	return CertAPI.IssueCert(auth, holder, serial, hash)

}
