package main

import (
	"flag"
	"goblockchain/wallet_server"
	"log"
)

func init() {
	log.SetPrefix("Wallet Server: ")
}

func main() {
	port := flag.Uint("port", 8080, "TCP Port Number for Wallet Server")
	gateway := flag.String("gateway", "http://127.0.0.1:9000", "Blockchain Gateway")
	flag.Parse()

	app := wallet_server.NewWalletServer(uint16(*port), *gateway)
	app.Run()
}