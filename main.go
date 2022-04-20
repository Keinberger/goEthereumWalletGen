package main

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// generating privateKey
	privateKey, err := crypto.GenerateKey()
	panicError(err)
	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println("PrivateKey")
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
	fmt.Println()

	// getting publicKey from privateKey
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panicError(errors.New("fetched publicKey does not have type of ecdsa.PublicKey"))
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	fmt.Println("PublicKey")
	fmt.Println(hexutil.Encode(publicKeyBytes)[2:])
	fmt.Println()

	// gettting address from publicKey
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address")
	fmt.Println(address)
}
