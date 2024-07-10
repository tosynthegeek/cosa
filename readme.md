# Cosa : 
**Cosa** is a blockchain built using Cosmos SDK & CometBFT, created with [Ignite CLI](https://ignite.com/cli),  featuring automated auction management with end-block processing for timely closure and updates.

## Overview

Cosa is a blockchain platform designed specifically for managing decentralized auctions. Built on the Cosmos SDK and using CometBFT for consensus, it provides a secure and transparent environment for creating, bidding on, and closing auctions.

## Features

- Create auctions with customizable parameters
- Place bids on active auctions
- Automated auction closure using end-block processing
- Real-time tracking of auction status and highest bids
- Secure and transparent transaction handling

## Prerequisites

- Go 1.16+
- Cosmos SDK
- Ignite CLI

## Installation

- ### Clone the repository: 
```
git@github.com:tosynthegeek/cosa.git
cd cosa
```

- ### Build the application: 
```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.
