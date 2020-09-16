module github.com/heystraightedge/straightedge

go 1.13

require (
	github.com/ChainSafe/go-schnorrkel v0.0.0-20200405005733-88cbf1b4c40d
	github.com/CosmWasm/wasmd v0.10.0
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/cosmos/go-bip39 v0.0.0-20200817134856-d632e0d11689
	github.com/gorilla/mux v1.7.4
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/sikkatech/go-substrate-bip39 v0.0.0-20191221093258-7f61eeaac83f
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.7
	github.com/tendermint/tm-db v0.5.1
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
)

replace github.com/cosmos/cosmos-sdk v0.39.1 => github.com/heystraightedge/cosmos-sdk v0.0.0-20200916052131-a4bcc34d2f53
