package main

import (
	"flag"
	g "myown/walletconfig"
)

var (
	user         = flag.String("username", "", "the username of the wallet")
	certPath     = flag.String("cert", "", "the path to the cert")
	keystorePath = flag.String("keystore", "", "the path to the keystore")
	mspName      = flag.String("msp_name", "", "the msp name")
)

func main() {
	g.WalletHandler(g.WalletMetadata{})
}
