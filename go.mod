module github.com/heystraightedge/straightedge

go 1.13

replace github.com/cosmwasm/wasmd => github.com/sikkatech/wasmd v0.0.0-20200123023242-b18141498392

require (
	github.com/ChainSafe/go-schnorrkel v0.0.0-20200102211924-4bcbc698314f
	github.com/btcsuite/btcd v0.0.0-20190824003749-130ea5bddde3 // indirect
	github.com/cosmos/cosmos-sdk v0.34.4-0.20200122195256-f18005d2f18b
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/cosmwasm/wasmd v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.7.3
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/sikkatech/go-substrate-bip39 v0.0.0-20191221093258-7f61eeaac83f
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.4.0
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.0
	github.com/tendermint/tm-db v0.4.0
	golang.org/x/crypto v0.0.0-20191219195013-becbf705a915
	golang.org/x/net v0.0.0-20190916140828-c8589233b77d // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)
