package main

import (
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func log(logKeys bool, msg string) {
	if logKeys {
		fmt.Println(msg)
	}
}

func main() {
	pwd := flag.String("pwd", "", "Password to encrypt keystore file")
	dir := flag.String("dir", "./wallets", "Directory to store keystore file")
	logKeys := flag.Bool("logging", false, "Dis/enable logging keys")
	flag.Parse()

	// generating privateKey
	privateKey, err := crypto.GenerateKey()
	panicError(err)
	privateKeyBytes := crypto.FromECDSA(privateKey)
	log(*logKeys, "PrivateKey\n"+hexutil.Encode(privateKeyBytes)[2:]+"\n")

	// getting publicKey from privateKey
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panicError(errors.New("fetched publicKey does not have type of ecdsa.PublicKey"))
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	log(*logKeys, "PublicKey\n"+hexutil.Encode(publicKeyBytes)[2:]+"\n")

	// generating keystore file from privateKey
	ks := keystore.NewKeyStore(*dir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, *pwd)
	panicError(err)

	fmt.Println("Ethereum Wallet (" + account.Address.Hex() + ") has been generated and stored in " + *dir)
}
