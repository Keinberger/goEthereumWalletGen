package main

import (
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
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
	mnemonic := flag.String("mnemonic", "", "Use a mnemonic phrase to generate the wallet")
	flag.Parse()

	// generating privateKey
	var privateKey *ecdsa.PrivateKey
	if *mnemonic != "" { // using mnemonic phrase
		wallet, err := hdwallet.NewFromMnemonic(strings.Trim(*mnemonic, "\""))
		panicError(err)
		account, err := wallet.Derive(hdwallet.DefaultBaseDerivationPath, false)
		panicError(err)
		privateKey, err = wallet.PrivateKey(account)
		panicError(err)
	} else {
		var err error
		privateKey, err = crypto.GenerateKey() // generating new random privateKey
		panicError(err)
	}

	// generating privateKey
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
	account, err := ks.ImportECDSA(privateKey, *pwd) // will throw error if one uses same mnemonic phrase again
	panicError(err)

	fmt.Println("Ethereum Wallet (" + account.Address.Hex() + ") has been generated and stored in " + *dir)
}
