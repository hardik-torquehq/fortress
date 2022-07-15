Fully functional blockchain with multi node consensus and fully funcitonal wallet engine built using GO libraries...feel free to fork and contribute to

Build Instructions

- Pretty much nothing...Currently its working for Go version 1.18.3 

Installation

- Run following commands
--  Start blockchain nodes with following commands:
--- cd blockchain_server
--- go run blockchain_server.go -port 9000
--- go run blockchain_server.go -port 9001
--- go run blockchain_server.go -port 9002
--- go run blockchain_server.go -port 9003

--  Start wallet engine with following commands from root folder:
--- go run main.go -port 8080 -gateway http://localhost:8080
--- go run main.go -port 8080 -gateway http://localhost:8081


Run 
- Go to browser and run http://localhost:8080 and in another tab http://localhost:8081
- Use mining keys to populate one of the wallets 
- Use blockchain address of other wallet as recipient and try sending some tokens and wolla!
- Check REST API status of blockchain nodes at http://localhost:9000/chains
  

Blockchain is runnable in production with few tweaks

Roadmap

Launch multinode testnet
faucets
explorer
EVM compatibility
