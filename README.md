<p align="center">
  <img src="./straightedge-logo.png" width="300">
</p>
<h3 align="center">A Cosmic Smart Contracting Platform.</h3>

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/heystraightedge/straightedge)](https://goreportcard.com/report/github.com/heystraightedge/straightedge)
[![API Reference](https://godoc.org/github.com/heystraightedge/straightedge?status.svg)](https://godoc.org/github.com/heystraightedge/straightedge)
[![Twitter Follow](https://img.shields.io/twitter/follow/heystraightedge.svg?label=Follow&style=social)](https://twitter.com/heystraightedge)

</div>

<div align="center">

### [Telegram](https://t.me/HeyStraightedge) | [Discord](https://discord.gg/rbamhbC) | [Medium](https://medium.com/straightedge)

### Participate in Straightedge!

</div>

Reference implementation of Straightedge. Built using the [comsos-sdk](https://github.com/cosmos/cosmos-sdk).

## Quick Start

```sh
make install
```

To join the mainnet, head over to the [mainnet launch repo](https://github.com/heystraightedge/mainnet).

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
