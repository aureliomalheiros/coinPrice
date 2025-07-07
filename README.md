# Coinprice CLI 💰

A lightweight cryptocurrency price checker written in Go that supports multiple coins and currencies using the CoinGecko API.

![coinprice](img/coinprice)

## 🌟 Features

- Get real-time prices for 30+ cryptocurrencies
- Supports both BRL (Brazilian Real) and USD (US Dollar)
- Simple JSON output option for integration with other tools
- Fast and lightweight (single binary)

## 📦 Installation

### Using Go (development)

```bash
go install github.com/aureliomalheiros/coinprice@latest
```

#### Pre-built binaries

Download from the [Releases page](https://github.com/aureliomalheiros/coinprice/releases)

### 🚀 Basic Usage

- Get Bitcoin price (default)
coinprice --brl       # Price in Brazilian Real
coinprice --usd       # Price in US Dollar

- Get specific cryptocurrency prices
coinprice --coin ethereum --brl
coinprice --coin solana --usd

- Multiple currencies
coinprice --brl --usd

- JSON output
coinprice --brl --json

- Options

```bash
-b, --brl        Get price in Brazilian Real
-u, --usd        Get price in US Dollar
-c, --coin string Specify cryptocurrency (default "bitcoin")
-j, --json       Output in JSON format
-h, --help       Show help
```

## Supported Cryptocurrencies 🪙

| Coin Name          | ID (for --coin flag)  |
|--------------------|-----------------------|
| Bitcoin            | `bitcoin`            |
| Ethereum           | `ethereum`           |
| Aave               | `aave`               |
| Solana             | `solana`             |
| Dogecoin           | `dogecoin`           |
| Litecoin           | `litecoin`           |
| Ripple (XRP)       | `ripple`             |
| Cardano            | `cardano`            |
| Polkadot           | `polkadot`           |
| Chainlink          | `chainlink`          |
| Uniswap            | `uniswap`            |
| Avalanche          | `avalanche`          |
| Cosmos             | `cosmos`             |
| Stellar            | `stellar`            |
| Monero             | `monero`             |
| Tron               | `tron`               |
| Algorand           | `algorand`           |
| VeChain            | `vechain`            |
| Filecoin           | `filecoin`           |
| Bitcoin Cash       | `bitcoin-cash`       |
| Dash               | `dash`               |
| Zcash              | `zcash`              |
| Tezos              | `tezos`              |
| Neo                | `neo`                |
| Elrond             | `elrond`             |
| Hedera             | `hedera`             |
| Fantom             | `fantom`             |
| Decentraland       | `decentraland`       |
| The Sandbox        | `the-sandbox`        |
| Axie Infinity      | `axie-infinity`      |
| Chiliz             | `chiliz`             |
| Enjin Coin         | `enjin-coin`         |
| Theta              | `theta`              |

> Note: All coin IDs are lowercase and use hyphens for multi-word names
