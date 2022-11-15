# Infrakit System Requirements

Infrakit will provide a toolkit to run two major systems of the Decentralized Infrastructure
Network: Infrastructure Providers, and Ingress Operators.

## Shared Requirements (KIT)

| Number  | Requirement |
| ------- | ----------- |
| [KIT-001](#kit-001) | Must support provisioning into a kubernetes cluster on the 3 major cloud providers (AWS, GCP, Azure). |
| [KIT-002](#kit-002) | Must be written in a popular programming language that encourages community contribution. |
| [KIT-003](#kit-003) | Must be developed as a permissively licensed OSS project. |

**Details/Notes:**

### KIT-003

The community is unlikely to have interest if licensed under anything other than `MIT`.

## Infrastructure Provider (PRO)

| Number  | Requirement |
| ------- | ----------- |
| [PRO-001](#pro-001) | Must support the Ethereum Layer 1 Protocol JSON-RPC API. |
| [PRO-002](#pro-002) | Must support provisioning new blockchain nodes quickly via filesystem snapshots. |
| [PRO-003](#pro-003) | Must authenticate and authorize requests signed with an Ethereum private key. |
| [PRO-004](#pro-004) | Must be configurable with an Ethereum private key to sign responses. |

**Details/Notes:**

### PRO-001

Including Trace API. Not including websockets or subscriptions. Mainnet and testnets.

### PRO-002

Quickly will vary based on the size of the chain, but using filesystem snapshots is generally
faster than all other sync options, especially for large historical chain data.

### PRO-003

- Signature used to authenticate the identity of the caller.
- Authorize usage against what the user has paid for.

## Ingress Operator (ING)

| Number  | Requirement |
| ------- | ----------- |
| [ING-001](#ing-001) | Must provide a consistent view of the blockchain. |
| [ING-002](#ing-002) | Must be configurable with an Ethereum private key to sign requests to infrastructure providers. |
| [ING-003](#ing-003) | Must be configurable to make use of multiple infrastructure providers. |

**Details/Notes:**

### ING-001

For example, consecutive call to `eth_blockNumber` should increase monotonically.

### ING-003

Sending requests to multiple providers can help ensure accuracy and availability.

- Load balancing, failover, cocktail, etc...