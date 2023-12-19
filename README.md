# bitcoin-puzzless

![Build Status](https://github.com/timchurchard/bitcoin-puzzles/workflows/Test/badge.svg)
![Coverage](https://img.shields.io/badge/Coverage-67.3%25-yellow)
[![License](https://img.shields.io/github/license/timchurchard/bitcoin-puzzles)](/LICENSE)
[![Release](https://img.shields.io/github/release/timchurchard/bitcoin-puzzles.svg)](https://github.com/timchurchard/bitcoin-puzzles/releases/latest)
[![GitHub Releases Stats of bitcoin-puzzles](https://img.shields.io/github/downloads/timchurchard/bitcoin-puzzles/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=timchurchard&repository=bitcoin-puzzles)

Simple golang implementation of some Bitcoin puzzles.

## Usage

## 32 Bitcoin

On the 15th January 2015 an interesting transaction appeared on the Bitcoin blockchain: [08389f34c98c...](https://www.blockchain.com/explorer/transactions/btc/08389f34c98c606322740c0be6a7125d9860bb8d5cb182c02f98461e5fa6cd15). This transaction sent 32 Bitcoin to 256 addresses. These addresses were made in a predictable way. Today many of these keys are found. 2^66 is not yet found [13zb1hQb...](https://www.blockchain.com/explorer/addresses/btc/13zb1hQbWVsc2S7ZTZnP2G4undNNpdh5so), this address now holds 6.6 Bitcoin. In April 2023 someone increased the prize to 1000 Bitcoin. More information can be found [here](https://privatekeys.pw/puzzles/bitcoin-puzzles-tx).

## Ballet Bobby (1 BTC)

On the 24th July 2020 Bobby Lee loaded two Ballet wallets with 1 BTC and part of the secret. He created a website [takebobbysbitcoin.com](https://www.takebobbysbitcoin.com/) with details.
