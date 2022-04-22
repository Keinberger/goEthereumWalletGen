# goEthereumWalletGen

## Requirements

Requires Go 1.18 or newer if run manually

## Usage

There are two options to use the script

* Running the script manually (requires Go preinstalled)
 `go run *.go -pwd passwordToEncrypt -dir ./directoryToStoreKeystore -mnemonic "insert mnemonic phrase" -logging`

* Make an executable `make build`
  * and running that afterwards `./main -pwd passwordToEncrypt -dir ./directoryToStoreKeystore -mnemonic "insert mnemonic phrase" -logging`

### Flags

There are several optional flags one can use

* `-pwd passwordToEncrypt`
  * specifies a password to encrypt the keystore file with
* `-dir ./directoryToStoreKeystore`
  * Specifies the direcotry to store the keystore file in
  * default is `./wallets`
* `-mnemonic` 
  * One may use a mnemonic key phrase to generate a wallet off of
* `-logging` `true` OR `false`
  * logs privateKey and publicKey into the terminal
  * default is false

## Important

Never share your keystore file or privateKey with anyone!
