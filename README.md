<p align="center">
  <img src="./straightedge-logo.png" width="300">
</p>
<h3 align="center">A Cosmic Smart Contracting Platform.</h3>

<div align="center">

[![API Reference](https://godoc.org/github.com/heystraightedge/straightedge?status.svg)](https://godoc.org/github.com/heystraightedge/straightedge)
[![Twitter Follow](https://img.shields.io/twitter/follow/heystraightedge.svg?label=Follow&style=social)](https://twitter.com/heystraightedge)
<!-- [![CircleCI](https://circleci.com/gh/heystraightedge/straightedge/tree/master.svg?style=shield)](https://circleci.com/gh/heystraightedge/straightedge/tree/master) -->
[![codecov](https://codecov.io/gh/heystraightedge/straightedge/branch/master/graph/badge.svg)](https://codecov.io/gh/heystraightedge/straightedge)
[![Go Report Card](https://goreportcard.com/badge/github.com/heystraightedge/straightedge)](https://goreportcard.com/report/github.com/heystraightedge/straightedge)
[![license](https://img.shields.io/github/license/heystraightedge/straightedge.svg)](https://github.com/heystraightedge/straightedge/blob/master/LICENSE)
[![LoC](https://tokei.rs/b1/github/heystraightedge/straightedge)](https://github.com/heystraightedge/straightedge)
<!-- [![GolangCI](https://golangci.com/badges/github.com/heystraightedge/straightedge.svg)](https://golangci.com/r/github.com/heystraightedge/straightedge) -->


</div>

<div align="center">

### [Telegram](https://t.me/HeyStraightedge) | [Discord](https://discord.gg/rbamhbC) | [Medium](https://medium.com/straightedge)

### Participate in Straightedge!

</div>

This repository hosts `strd`, the reference implementation of the Straightedge blockchain.  It is built using the [comsos-sdk](https://github.com/cosmos/cosmos-sdk) and supports the [CosmWasm](https://github.com/CosmWasm/wasmd) smart contracting system. 

## Quick Start

```sh
make install
make test
```
if you are using a linux without X or headless linux, look at [this article](https://ahelpme.com/linux/dbusexception-could-not-get-owner-of-name-org-freedesktop-secrets-no-such-name) or [#31](https://github.com/cosmwasm/wasmd/issues/31#issuecomment-577058321).

<!-- To set up a single node testnet, [look at the deployment documentation](./docs/deploy-testnet.md). -->

<!-- If you want to deploy a whole cluster, [look at the network scripts](./networks/README.md). -->

To join the Straightedge mainnet, head over to the [mainnet launch repo](https://github.com/heystraightedge/mainnet).

## Importing Lockdrop Keys

During the lockdrop, participants submitted SR25519 keys to be creditted in the Edgeware/Straightedge networks.  To support these keys we have added support to the Straightedge CLI wallet for SR25519 keys. To add your lockdrop keys, please use the following commands:

```
strcli keys add [name] --algo sr25519 --recover
[insert mnemonic here]
```

As sr25519 is a relatively new standard, it is not supported by many wallets as of now.  It is recommended to move funds to a new secp256k1 address, for easier use in Cosmos wallets.

<!-- 
## License

Licensed under the [Apache v2 License](LICENSE.md). -->