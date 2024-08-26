module github.com/ontio/ontology-crypto

go 1.12

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0 // indirect
	github.com/ethereum/go-ethereum v1.13.15
	github.com/itchyny/base58-go v0.1.0
	github.com/stretchr/testify v1.8.4
	golang.org/x/crypto v0.26.0
)

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191029031824-8986dd9e96cf
