package gofs

//go:generate abigen --lang go -pkg gofs --abi contracts/IGOFS.abi --bin contracts/IGOFS.bin --type IGOFS --out contract_igofs.go
//go:generate abigen --lang go -pkg gofs --abi contracts/GOFS.abi --bin contracts/GOFS.bin --type GOFS --out contract_gofs.go
